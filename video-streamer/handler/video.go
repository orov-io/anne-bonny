package handler

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
)

type Video struct {
	storageServiceHost *url.URL
}

const videoPath = "/video"

func NewVideoHandler(storageServiceHost *url.URL) *Video {
	return &Video{
		storageServiceHost: storageServiceHost,
	}
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
	fmt.Println("adding video")
	e.GET(videoPath, v.GetVideoHandler)
}
