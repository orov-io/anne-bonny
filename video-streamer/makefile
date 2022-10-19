DOKKU_ENV_FILE=dokku.env
DOKKU_ENVS=`tr '\n' ' ' < $(DOKKU_ENV_FILE)`

run:
	@echo "Running air as hot-reload provider"
	@air

docker-build:
	@echo "Building ts-video-streamer:latest docker image"
	@docker build -t orovium/ts-video-streamer .

docker-run: docker-stop
	@echo "Running ts-video-streamer image as video-streamer container on port 4001"
	@echo "We carefully recommend to run the image using docker-compose from the root project."
	@docker run --env PORT=80 -d -p 4001:80 --name video-streamer orovium/ts-video-streamer

docker-stop:
	@echo "stopping video-streamer containter"
	-@docker stop video-streamer
	-@docker rm video-streamer

docker-push:
	@echo "Pushing to docker hub repo"
	@docker push orovium/ts-video-streamer:latest
	@docker tag orovium/ts-video-streamer orovium/ts-video-streamer:$(shell git rev-parse HEAD)
	@docker push orovium/ts-video-streamer:$(shell git rev-parse HEAD)

docker-enter:
	@echo "Attaching a tty to video-streamer container"
	@docker exec -it video-streamer /bin/ash

dokku-deploy:
	@echo "Setting env vars"
	@ssh dokku@anne-bonny.dokku config:set --no-restart video-streamer $(DOKKU_ENVS)
	@echo "Deploying to dokku service"
	@ssh dokku@anne-bonny.dokku git:from-image video-streamer orovium/ts-video-streamer:$(shell git rev-parse HEAD)

tidy:
	@echo "Installing project dependencies"
	@go mod tidy

prepare: tidy
	@echo "Installing air to hot-reloading"
	@go install github.com/cosmtrek/air@latest
