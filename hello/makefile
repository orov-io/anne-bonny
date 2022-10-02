run:
	@echo "Running air as hot-reload provider"
	@PORT=4000 air

docker-build:
	@echo "Building hello:latest docker image"
	@docker build -t anne-bonny/hello .

docker-run: docker-stop
	@echo "Running hello container as hello"
	@docker run --env PORT=1323 -d -p 4000:1323 --name hello anne-bonny/hello

docker-stop:
	@echo "stopping hello containter"
	-@docker stop hello
	@echo "removing hello containter"
	-@docker rm hello

docker-push:
	@echo "Pushing to private repo"
	@docker push anne-bonny/hello:latest

docker-enter:
	@echo "Attaching a tty to hello container"
	@docker exec -it hello /bin/ash
tidy:
	@echo "Installing project dependencies"
	@go mod tidy

prepare: tidy
	@echo "Installing air to hot-reloading"
	@go install github.com/cosmtrek/air@latest
