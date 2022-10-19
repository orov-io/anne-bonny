run:
	@echo "Running air as hot-reload provider"
	@air

docker-build:
	@echo "Building storage-azure:latest docker image"
	@docker build -t orovio.azurecr.io/storage-azure .

docker-run: docker-stop
	@echo "Running storage-azure container as storage-azure"
	@docker run --env PORT=1323 -d -p 1323:1323 --name storage-azure orovio.azurecr.io/storage-azure

docker-stop:
	@echo "stopping storage-azure containter"
	-@docker stop storage-azure
	-@docker rm storage-azure

docker-push:
	@echo "Pushing to private repo"
	@docker push orovio.azurecr.io/storage-azure:latest

docker-enter:
	@echo "Attaching a tty to storage-azure container"
	@docker exec -it storage-azure /bin/ash
tidy:
	@echo "Installing project dependencies"
	@go mod tidy

prepare: tidy
	@echo "Installing air to hot-reloading"
	@go install github.com/cosmtrek/air@latest
