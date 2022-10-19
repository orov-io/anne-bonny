package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/orov-io/anne-bonny/storage/azure/handler"
	"github.com/orov-io/maryread"
	mrHandler "github.com/orov-io/maryread/handler"
)

const portEnvKey = "PORT"
const storageAccountNameEnvKey = "STORAGE_ACCOUNT_NAME"
const storageAccessKeyEnvKey = "STORAGE_ACCESS_KEY"
const storageContainerEnvKey = "STORAGE_CONTAINER"

var port string
var storageAccountName string
var storageAccessKey string
var storageContainer string

func init() {
	parseEnvs()
}

func parseEnvs() {
	parsePort()
	parseStorageCredentials()
}

func parsePort() {
	parsedPort, ok := os.LookupEnv(portEnvKey)
	if !ok {
		panic(fmt.Sprintf("Please specify the port number for the HTTP server with the environmnet variable %v", portEnvKey))
	}

	_, err := strconv.Atoi(parsedPort)
	if err != nil {
		panic(fmt.Sprintf("%v env variable must be an integer", portEnvKey))
	}

	port = fmt.Sprintf(":%v", parsedPort)
}

func parseStorageCredentials() {
	var ok bool
	storageAccountName, ok = os.LookupEnv(storageAccountNameEnvKey)
	if !ok {
		panic(fmt.Sprintf("Please specify the name of an Azure storage account in environment variable %v", storageAccountNameEnvKey))
	}

	storageAccessKey, ok = os.LookupEnv(storageAccessKeyEnvKey)
	if !ok {
		panic(fmt.Sprintf("Please specify the access key to an Azure storage account in environment variable %v", storageAccessKeyEnvKey))
	}

	storageContainer, ok = os.LookupEnv(storageContainerEnvKey)
	if !ok {
		panic(fmt.Sprintf("Please specify the container inside the Azure storage account in environment variable %v", storageContainerEnvKey))
	}
}

func main() {
	app := maryread.Default()
	addHandlers(app)
	initApp(app)
}

func addHandlers(app *maryread.App) {
	videoHandler, err := handler.NewVideoHandler(storageAccountName, storageAccessKey, storageContainer)
	if err != nil {
		panic(fmt.Sprintf("Unable to start video handler due to: %v", err))
	}
	videoHandler.AddHandlers(app.Router())

	mrHandler.NewPingHandler().AddHandlers(app.Router())
}

func initApp(app *maryread.App) {
	app.Router().Logger.Fatal(app.Router().Start(port))
}
