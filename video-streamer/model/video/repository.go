package video

type Repository interface {
	GetVideo(id string) (*Video, error)
}
