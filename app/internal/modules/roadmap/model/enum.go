package roadmapmodel

type Status string

const (
	StatusCompleted  Status = "completed"
	StatusInProgress Status = "in-progress"
	StatusComingSoon Status = "coming-soon"
)

func (s Status) IsValid() bool {
	switch s {
	case StatusCompleted, StatusInProgress, StatusComingSoon:
		return true
	}
	return false
}
