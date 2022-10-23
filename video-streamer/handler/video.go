package handler

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	m "github.com/orov-io/anne-bonny/video-streamer/middleware"
	"github.com/orov-io/anne-bonny/video-streamer/model/video"
	"github.com/orov-io/maryread"
)

type Video struct {
	// This will disappear when we code the private sdk and inject it in the context.
	storageServiceHost *url.URL
}

const videoPath = "/video"

func NewVideoHandler(storageServiceHost *url.URL) *Video {
	return &Video{
		storageServiceHost: storageServiceHost,
	}
}

type getVideoQueryRequest struct {
	Id string `query:"id" validate:"required,uuid"`
}

func (v *Video) GetVideoHandler(c echo.Context) error {
	query := new(getVideoQueryRequest)
	err := maryread.Bindlidate(c, query)
	if err != nil {
		return err
	}

	currentVideo, err := m.MustGetFactory(c).Video.GetVideo(query.Id)

	if err != nil {
		if errors.Is(err, video.ErrVideoNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return err
	}

	storageVideoURI := v.storageServiceHost.JoinPath("/storage/").JoinPath(currentVideo.Path)

	resp, err := http.Get(storageVideoURI.String())
	if err != nil {
		return err
	}

	return c.Stream(resp.StatusCode, resp.Header.Get("Content-Type"), resp.Body)
}

func (v *Video) AddHandlers(e *echo.Echo) {
	e.GET(videoPath, v.GetVideoHandler)
}
