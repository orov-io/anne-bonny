package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"

	_ "github.com/lib/pq" // Here we initializes the database
	"github.com/orov-io/anne-bonny/video-streamer/handler"
	"github.com/orov-io/maryread"
	mrHandler "github.com/orov-io/maryread/handler"
	"github.com/orov-io/maryread/middleware"
)

const portEnvKey = "PORT"
const storageServiceHostEnvKey = "STORAGE_SERVICE_HOST"

var port string
var storageServiceHost *url.URL

func init() {
	parseEnvs()
}

func parseEnvs() {
	parsePort()
	parseStorageServiceHost()
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

func parseStorageServiceHost() {
	storageServiceHostRaw, ok := os.LookupEnv(storageServiceHostEnvKey)
	if !ok {
		panic(fmt.Sprintf("Please specify the host for the storage service with the environment variable %v", storageServiceHostEnvKey))
	}

	var err error
	storageServiceHost, err = url.Parse(storageServiceHostRaw)
	if err != nil {
		panic(fmt.Sprintf("Provided storage service host <%v> is not a valid URI. Error parsing: %v", storageServiceHostRaw, err))
	}
}

func main() {
	app := maryread.Default()
	// To use this, you must import the pq driver as:  _ "github.com/lib/pq"
	// TODO: Put this in this README and in the maryread README
	app.Router().Use(middleware.SQLX())
	addHandlers(app)
	initApp(app)
}

func addHandlers(app *maryread.App) {
	handler.NewVideoHandler(storageServiceHost).AddHandlers(app.Router())
	mrHandler.NewPingHandler().AddHandlers(app.Router())
}

func initApp(app *maryread.App) {
	app.Router().Logger.Fatal(app.Router().Start(port))
}
