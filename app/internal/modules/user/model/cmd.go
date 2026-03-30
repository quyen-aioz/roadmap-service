package usermodel

import (
	"fmt"
	"strings"
)

type FindQueryBuilder struct {
	ID       string
	Username string
}

func (q *FindQueryBuilder) Build() (string, []any, error) {
	var clauses []string
	var args []any

	if q.ID != "" {
		clauses = append(clauses, "id = ?")
		args = append(args, q.ID)
	}

	if q.Username != "" {
		clauses = append(clauses, "username = ?")
		args = append(args, q.Username)
	}

	if len(clauses) == 0 {
		return "", nil, fmt.Errorf("empty query")
	}

	whereClause := strings.Join(clauses, " AND ")

	return whereClause, args, nil
}
