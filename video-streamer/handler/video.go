package handler

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
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

type GetVideoRequest struct {
}

func (v *Video) GetVideoHandler(c echo.Context) error {
	storageVideoURI := v.storageServiceHost.JoinPath("/storage/SampleVideo_1280x720_1mb.mp4")

	resp, err := http.Get(storageVideoURI.String())
	if err != nil {
		return err
	}

	return c.Stream(resp.StatusCode, resp.Header.Get("Content-Type"), resp.Body)
}

func (v *Video) AddHandlers(e *echo.Echo) {
	e.GET(videoPath, v.GetVideoHandler)
}
