package roadmapmodel

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

type FindQueryBuilder struct {
	ID      string
	Title   string
	Status  Status
	GroupID string
}

func (q *FindQueryBuilder) Build() (string, []any) {
	var clauses []string
	var args []any

	clauses = append(clauses, "deleted_at IS NULL")

	if q.ID != "" {
		clauses = append(clauses, "id = ?")
		args = append(args, q.ID)
	}

	if q.Title != "" {
		clauses = append(clauses, "title LIKE ?")
		args = append(args, "%"+q.Title+"%")
	}

	if q.Status != "" {
		clauses = append(clauses, "status = ?")
		args = append(args, q.Status)
	}

	if q.GroupID != "" {
		clauses = append(clauses, "group_id = ?")
		args = append(args, q.GroupID)
	}

	// Join all clauses with " AND "
	whereClause := " WHERE " + strings.Join(clauses, " AND ")

	return whereClause, args
}

type RoadmapUpdateBuilder struct {
	Title         *string
	Content       *string
	Status        *Status
	GroupID       *GroupID
	CTALabel      *string
	CTALink       *string
	StartDate     *time.Time
	EndDate       *time.Time
	ThumbnailURL  *string
	ThumbnailType *string
}

func (u *RoadmapUpdateBuilder) Build() (string, []any, error) {
	var clauses []string
	var args []any

	if u.Title != nil {
		clauses = append(clauses, "title = ?")
		args = append(args, *u.Title)
	}
	if u.Content != nil {
		clauses = append(clauses, "content = ?")
		args = append(args, *u.Content)
	}
	if u.Status != nil {
		clauses = append(clauses, "status = ?")
		args = append(args, *u.Status)
	}
	if u.GroupID != nil {
		clauses = append(clauses, "group_id = ?")
		args = append(args, *u.GroupID)
	}
	if u.CTALabel != nil {
		clauses = append(clauses, "cta_label = ?")
		args = append(args, *u.CTALabel)
	}
	if u.CTALink != nil {
		clauses = append(clauses, "cta_link = ?")
		args = append(args, *u.CTALink)
	}
	if u.StartDate != nil {
		clauses = append(clauses, "start_date = ?")
		args = append(args, u.StartDate.UTC())
	}
	if u.EndDate != nil {
		clauses = append(clauses, "end_date = ?")
		args = append(args, u.EndDate.UTC())
	}
	if u.ThumbnailURL != nil {
		clauses = append(clauses, "thumbnail_url = ?")
		args = append(args, *u.ThumbnailURL)
	}
	if u.ThumbnailType != nil {
		clauses = append(clauses, "thumbnail_type = ?")
		args = append(args, *u.ThumbnailType)
	}

	if len(clauses) == 0 {
		return "", nil, fmt.Errorf("empty update")
	}

	clauses = append(clauses, "updated_at = ?")

	args = append(args, time.Now())

	setClause := " SET " + strings.Join(clauses, ", ")

	return setClause, args, nil
}

func GenerateHexID() string {
	b := make([]byte, 8) // 8 bytes = 16 hex characters
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}
