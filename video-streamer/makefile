DOKKU_ENV_FILE=dokku.env
DOKKU_ENVS=`tr '\n' ' ' < $(DOKKU_ENV_FILE)`

docker-build:
	@echo "Building ab-video-streamer:latest docker image"
	@docker build -t orovium/ab-video-streamer .

docker-push:
	@echo "Pushing to docker hub repo"
	@docker push orovium/ab-video-streamer:latest
	@docker tag orovium/ab-video-streamer orovium/ab-video-streamer:$(shell git rev-parse HEAD)
	@docker push orovium/ab-video-streamer:$(shell git rev-parse HEAD)

dokku-deploy:
	@echo "Setting env vars"
	-@ssh dokku@anne-bonny.dokku config:set --no-restart ab-video-streamer $(DOKKU_ENVS)
	@echo "Deploying to dokku service"
	@ssh dokku@anne-bonny.dokku git:from-image ab-video-streamer orovium/ab-video-streamer:$(shell git rev-parse HEAD)

tidy:
	@echo "Installing project dependencies"
	@go mod tidy

prepare: tidy
	@echo "Installing air to hot-reloading"
	@go install github.com/cosmtrek/air@latest

migrate-up:
	@echo "Applying pending migrations"
	@goose -dir ./migration up

migrate-up-one:
	@echo "Applying next pending mitration"
	@goose -dir ./migration up-by-one

migrate-down:
	@echo "Reverting lastest migration"
	@goose -dir ./migration down

migrate-reset:
	@echo "Reverting all migrations"
	@goose -dir ./migration reset

migrate-redo:
	@echo "Re-run lastest migration"
	@goose -dir ./migration redo

migrate-status:
	@echo "Listing migration status"
	@goose -dir ./migration status