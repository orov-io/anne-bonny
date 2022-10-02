package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/orov-io/TS-streamer/video-streamer/handler"
	"github.com/orov-io/maryread"
	mrHandler "github.com/orov-io/maryread/handler"
)

const portEnvKey = "PORT"

var port string

func init() {
	parseEnvs()
}

func parseEnvs() {
	parsePort()
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

func main() {
	app := maryread.Default()
	addHandlers(app)
	initApp(app)
}

func addHandlers(app *maryread.App) {
	handler.NewHelloHandler().AddHandlers(app.Router())
	mrHandler.NewPingHandler().AddHandlers(app.Router())
}

func initApp(app *maryread.App) {
	app.Router().Logger.Fatal(app.Router().Start(port))
}
