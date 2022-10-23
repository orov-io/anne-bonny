package video

import "github.com/google/uuid"

type Video struct {
	ID   uuid.UUID `json:"id"`
	Path string    `json:"path"`
}
