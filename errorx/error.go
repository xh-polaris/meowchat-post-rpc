package errorx

import "google.golang.org/grpc/status"

var (
	ErrNoSuchPost      = status.Error(10301, "no such post")
	ErrInvalidObjectId = status.Error(10302, "invalid objectId")
)
