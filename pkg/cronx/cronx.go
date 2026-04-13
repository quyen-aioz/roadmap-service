package cronx

import (
	"context"
	"log"
	"sync"
	"time"
)

type Job interface {
	Name() string
	Run(ctx context.Context) error
}

type Option func(*entry)

// RunOnStart makes the job run immediately when Start() is called,
// instead of waiting for the first tick.
func RunOnStart() Option {
	return func(e *entry) {
		e.runOnStart = true
	}
}

type entry struct {
	job        Job
	interval   time.Duration
	runOnStart bool
}

type Runner struct {
	entries []entry
	wg      sync.WaitGroup

	// mu protects entries from concurrent Register calls
	// before Start() is called
	mu sync.Mutex
}

func New() *Runner {
	return &Runner{}
}

func (r *Runner) Register(job Job, interval time.Duration, opts ...Option) {
	r.mu.Lock()
	defer r.mu.Unlock()

	e := entry{
		job:      job,
		interval: interval,
	}

	for _, opt := range opts {
		opt(&e)
	}

	r.entries = append(r.entries, e)
}

// Start spawns a goroutine for each registered job.
// It is non-blocking — it returns immediately after spawning.
// Pass the app-level context so all jobs stop on shutdown.
func (r *Runner) Start(ctx context.Context) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, e := range r.entries {
		e := e // capture loop variable — important in Go < 1.22
		r.wg.Add(1)
		go r.run(ctx, e)
	}

	log.Printf("[cronx] started %d job(s)", len(r.entries))
}

func (r *Runner) Stop() {
	r.wg.Wait()
	log.Println("[cronx] all jobs stopped")
}

func (r *Runner) run(ctx context.Context, e entry) {
	defer r.wg.Done()

	log.Printf("[cronx] job %q registered, interval: %s", e.job.Name(), e.interval)

	// run immediately on start if opted in
	if e.runOnStart {
		r.execute(ctx, e.job)
	}

	ticker := time.NewTicker(e.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			r.execute(ctx, e.job)

		case <-ctx.Done():
			log.Printf("[cronx] job %q stopping", e.job.Name())
			return
		}
	}
}

// execute runs a single job invocation with panic recovery.
// A panic in job.Run() is caught and logged — it never crashes the server.
func (r *Runner) execute(ctx context.Context, job Job) {
	// inner func so defer recover() scopes to THIS execution only
	// not the entire goroutine loop
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("[cronx] job %q panic recovered: %v", job.Name(), rec)
			}
		}()

		log.Printf("[cronx] job %q running", job.Name())

		if err := job.Run(ctx); err != nil {
			log.Printf("[cronx] job %q error: %v", job.Name(), err)
			return
		}

		log.Printf("[cronx] job %q completed", job.Name())
	}()
}
