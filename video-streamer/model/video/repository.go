package videoModel

type VideoRepository interface {
	GetVideo(id string) (*Video, error)
}
