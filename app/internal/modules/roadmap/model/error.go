package roadmapmodel

import "roadmap/pkg/apperror"

var ErrInvalidStatus = apperror.New(apperror.ErrInvalidStatus, "invalid status")
var ErrInvalidGroupID = apperror.New(apperror.ErrInvalidGroupID, "invalid group id")
var ErrRoadmapNotFound = apperror.New(apperror.ErrRoadmapNotFound, "roadmap not found")
var ErrUnsupportedMediaType = apperror.New(apperror.ErrUnsupportedMediaType, "unsupported media type")
