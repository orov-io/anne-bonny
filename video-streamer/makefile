DOKKU_ENV_FILE=dokku.env
DOKKU_ENVS=`tr '\n' ' ' < $(DOKKU_ENV_FILE)`

run:
	@echo "Running air as hot-reload provider"
	@air

docker-build:
	@echo "Building ab-video-streamer:latest docker image"
	@docker build -t orovium/ab-video-streamer .

docker-run: docker-stop
	@echo "Running ab-video-streamer container on port 4000"
	@echo "We carefully recommend to run the image using docker-compose from the root project."
	@docker run --env PORT=80 -d -p 4000:80 --name video-streamer orovium/ab-video-streamer

docker-stop:
	@echo "stopping video-streamer containter"
	-@docker stop video-streamer
	-@docker rm video-streamer

docker-push:
	@echo "Pushing to docker hub repo"
	@docker push orovium/ab-video-streamer:latest
	@docker tag orovium/ab-video-streamer orovium/ab-video-streamer:$(shell git rev-parse HEAD)
	@docker push orovium/ab-video-streamer:$(shell git rev-parse HEAD)

docker-enter:
	@echo "Attaching a tty to video-streamer container"
	@docker exec -it video-streamer /bin/ash

dokku-deploy:
	@echo "Setting env vars"
	@ssh dokku@anne-bonny.dokku config:set --no-restart ab-video-streamer $(DOKKU_ENVS)
	@echo "Deploying to dokku service"
	@ssh dokku@anne-bonny.dokku git:from-image ab-video-streamer orovium/ab-video-streamer:$(shell git rev-parse HEAD)

tidy:
	@echo "Installing project dependencies"
	@go mod tidy

prepare: tidy
	@echo "Installing air to hot-reloading"
	@go install github.com/cosmtrek/air@latest
