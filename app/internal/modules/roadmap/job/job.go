package roadmapjob

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	roadmaprepo "roadmap/app/internal/modules/roadmap/repo"
	"roadmap/pkg/aws3sx"
)

// CleanupOrphanThumbnailsJob deletes files in AIOZ storage that are no longer
// referenced by any roadmap task in the DB.
// It only deletes files older than the grace period to avoid deleting
// files that were just uploaded but not yet saved.
type CleanupOrphanThumbnailsJob struct {
	repo         roadmaprepo.Repository
	s3Client     *aws3sx.Client
	gracePeriod  time.Duration
	objectPrefix string
}

func NewCleanupOrphanThumbnailsJob(
	repo roadmaprepo.Repository,
	s3Client *aws3sx.Client,
	objectPrefix string,
) *CleanupOrphanThumbnailsJob {
	return &CleanupOrphanThumbnailsJob{
		repo:         repo,
		s3Client:     s3Client,
		gracePeriod:  24 * time.Hour, // safe default: 48h
		objectPrefix: strings.Trim(objectPrefix, "/"),
	}
}

func (j *CleanupOrphanThumbnailsJob) Name() string {
	return "cleanup_thumbnails"
}

func (j *CleanupOrphanThumbnailsJob) Run(ctx context.Context) error {
	objects, err := j.s3Client.ListObjects(ctx, aws3sx.ListObjectsInput{
		Prefix: j.objectPrefix,
	})
	if err != nil {
		return fmt.Errorf("list objects: %w", err)
	}

	if len(objects) == 0 {
		log.Printf("[%s] no objects found in storage, skipping", j.Name())
		return nil
	}

	log.Printf("[%s] found %d object(s) in storage", j.Name(), len(objects))

	// get all thumbnail_urls currently in use from DB
	activeURLs, err := j.repo.GetAllThumbnailURLs(ctx)
	if err != nil {
		return fmt.Errorf("get active thumbnail urls: %w", err)
	}

	// build a set for O(1) lookup
	// key: s3 object key (extracted from full public URL)
	activeKeys := make(map[string]struct{}, len(activeURLs))
	var invalidURLCount int
	for _, url := range activeURLs {
		if url == "" {
			continue
		}
		key := j.s3Client.ExtractKeyFromURL(url)
		if key == "" {
			invalidURLCount++
			log.Printf("[%s] invalid thumbnail url format, skipping match: %q", j.Name(), url)
			continue
		}
		activeKeys[key] = struct{}{}
	}

	if invalidURLCount > 0 {
		return fmt.Errorf("aborting cleanup: %d thumbnail url(s) could not be parsed into object keys", invalidURLCount)
	}

	log.Printf("[%s] found %d active thumbnail(s) in DB", j.Name(), len(activeKeys))

	// find orphans: in storage but NOT in DB, AND older than grace period
	now := time.Now()
	var orphanKeys []string

	for _, obj := range objects {
		// skip if still referenced in DB
		if _, isActive := activeKeys[obj.Key]; isActive {
			continue
		}

		// skip if younger than grace period
		// (admin might have uploaded but not saved yet)
		age := now.Sub(obj.LastModified)
		if age < j.gracePeriod {
			log.Printf("[%s] skipping %q — age %s is within grace period", j.Name(), obj.Key, age.Round(time.Minute))
			continue
		}

		orphanKeys = append(orphanKeys, obj.Key)
	}

	if len(orphanKeys) == 0 {
		log.Printf("[%s] no orphans found, nothing to clean up", j.Name())
		return nil
	}

	log.Printf("[%s] found %d orphan(s) to delete", j.Name(), len(orphanKeys))

	// delete orphans one by one, log errors but don't stop on failure
	// we don't want one bad delete to abort the rest
	var deletedCount, failedCount int
	for _, key := range orphanKeys {
		if err := j.s3Client.DeleteObject(ctx, key); err != nil {
			log.Printf("[%s] failed to delete %q: %v", j.Name(), key, err)
			failedCount++
			continue
		}
		log.Printf("[%s] deleted orphan: %q", j.Name(), key)
		deletedCount++
	}

	log.Printf("[%s] cleanup done — deleted: %d, failed: %d", j.Name(), deletedCount, failedCount)

	if failedCount > 0 {
		return fmt.Errorf("%d deletion(s) failed out of %d orphans", failedCount, len(orphanKeys))
	}

	return nil
}
