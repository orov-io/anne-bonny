package video

import "errors"

var ErrVideoNotFound = errors.New("Video not found in db")

type Repository interface {
	GetVideo(id string) (*Video, error)
}
