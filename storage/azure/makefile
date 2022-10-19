run:
	@echo "Running air as hot-reload provider"
	@air

docker-build:
	@echo "Building ab-storage-azure:latest docker image"
	@docker build -t orovium/ab-storage-azure .

docker-run: docker-stop
	@echo "Running ab-storage-azure container on port 4001"
	@docker run --env PORT=80 -d -p 4001:80 --name ab-storage-azure orovium/ab-storage-azure

docker-stop:
	@echo "stopping ab-storage-azure containter"
	-@docker stop ab-storage-azure
	-@docker rm ab-storage-azure

docker-push:
	@echo "Pushing to docker hub repo"
	@docker push orovium/ab-storage-azure:latest
	@docker tag orovium/ab-storage-azure orovium/ab-storage-azure:$(shell git rev-parse HEAD)
	@docker push orovium/ab-storage-azurer:$(shell git rev-parse HEAD)

docker-enter:
	@echo "Attaching a tty to storage-azure container"
	@docker exec -it ab-storage-azure /bin/ash

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
