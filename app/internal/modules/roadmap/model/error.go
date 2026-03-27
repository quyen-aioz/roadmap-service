package roadmapmodel

import "roadmap/pkg/apperror"

var ErrInvalidStatus = apperror.New(apperror.ErrInvalidStatus, "invalid status")
var ErrRoadmapNotFound = apperror.New(apperror.ErrNotFound, "roadmap not found")
