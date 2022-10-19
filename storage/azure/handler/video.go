package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/labstack/echo/v4"
	"github.com/orov-io/maryread"
)

type Video struct {
	storageAccountName string
	storageAccesskey   string
	storageContainer   string
	credential         *azblob.SharedKeyCredential
	client             *azblob.Client
}

const videoPath = "/storage/:videoName"

// Bellow consts will be used to check the provided storage variables:
const (
	existBlobFile = "exist.txt"
	destFileName  = "test_download_file.*.txt"
	destFileDir   = "./"
)

func NewVideoHandler(storageAccountName, storageAccesskey, storageContainer string) (*Video, error) {
	video := &Video{
		storageAccountName: storageAccountName,
		storageAccesskey:   storageAccesskey,
		storageContainer:   storageContainer,
	}
	err := video.getStorage()
	if err != nil {
		return nil, err
	}
	err = video.checkCredentials()
	if err != nil {
		return nil, err
	}

	return video, nil
}

func (v *Video) getStorage() error {
	var err error
	v.credential, err = azblob.NewSharedKeyCredential(v.storageAccountName, v.storageAccesskey)
	if err != nil {
		return err
	}

	v.client, err = azblob.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", v.storageAccountName), v.credential, nil)
	if err != nil {
		return err
	}
	return nil
}

/*
checkCredentials is used to check the integrity of the provided storage variables.

As azure not provide a clear way to know if an account is valid (you can try
to download from a not valid account and it only will fails with the default timeout...)
we need to serch for a sample file to initialize the service.
*/
func (v *Video) checkCredentials() error {
	destFile, err := os.CreateTemp(destFileDir, destFileName)
	if err != nil {
		return fmt.Errorf("unable to create file due to: %v", err)
	}
	defer os.Remove(destFileDir + destFile.Name())
	ctxt, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	_, err = v.client.DownloadFile(ctxt, v.storageContainer, existBlobFile, destFile, nil)

	if err != nil {
		return fmt.Errorf("unable to fetch files from %v/%v due to %v. Please, assert that provided container in provided account exists, and it have a file called %s in the root folder", v.storageAccountName, v.storageContainer, err, existBlobFile)
	}
	return nil
}

type getVideoRequest struct {
	VideoName string `param:"videoName"`
}

// GetVideoHandler tries to obtain the video from provided azure account:container
func (v *Video) GetVideoHandler(c echo.Context) error {
	logger := maryread.GetLogger(c)
	var request getVideoRequest
	err := c.Bind(&request)
	if err != nil {
		return err
	}

	stream, err := v.client.DownloadStream(context.Background(), v.storageContainer, request.VideoName, nil)
	if err != nil {
		switch v := err.(type) {
		case *azcore.ResponseError:
			logger.Debug().Msgf("Unable to find blob due to: %v", err.Error())
			return echo.NewHTTPError(v.StatusCode)
		}
		logger.Warn().Msgf("Not expected error searching video %s: %v", request.VideoName, err.Error())
		return err
	}

	return c.Stream(http.StatusOK, "video/mp4", stream.Body)
}

func (v *Video) AddHandlers(e *echo.Echo) {
	e.GET(videoPath, v.GetVideoHandler)
}
