docker-build:
	@echo "Building ab-storage-azure:latest docker image"
	@docker build -t orovium/ab-storage-azure .

docker-push:
	@echo "Pushing to docker hub repo"
	@docker push orovium/ab-storage-azure:latest
	@docker tag orovium/ab-storage-azure orovium/ab-storage-azure:$(shell git rev-parse HEAD)
	@docker push orovium/ab-storage-azurer:$(shell git rev-parse HEAD)

dokku-deploy:
	@echo "Setting env vars"
	-@ssh dokku@anne-bonny.dokku config:set --no-restart ab-video-streamer $(DOKKU_ENVS)
	@echo "Deploying to dokku service"
	@ssh dokku@anne-bonny.dokku git:from-image ab-video-streamer orovium/ab-video-streamer:$(shell git rev-parse HEAD)

tidy:
	@echo "Installing project dependencies"
	@go mod tidy
