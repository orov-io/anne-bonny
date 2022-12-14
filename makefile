include testing.env
export $(shell sed 's/=.*//' testing.env)

include local.env
export $(shell sed 's/=.*//' local.env)

include override.env
export $(shell sed 's/=.*//' override.env)

up: build
	@echo "Booting anne-bonny services with docker-compose"
	@docker-compose up -d

build:
	@echo "Building anne-bonny services"
	@docker-compose build --parallel

logs:
	@echo "Tailing anne-bonny containers logs"
	@docker-compose logs -f

down:
	@echo "Shutting down all anne-bonny containers"
	-@docker-compose down

reboot: down up logs

up-tool:
	@echo "Booting anne-bonny tools with docker compose"
	@docker-compose -f docker-compose.tool.yml up -d

# tool is a shortcut of up-tool so you can type make up tool logs :)
tool: up-tool

build-tool:
	@echo "Building anne-bonny services"
	@docker-compose -f docker-compose.tool.yml build --parallel

logs-tool:
	@echo "Tailing anne-bonny containers logs"
	@docker-compose -f docker-compose.tool.yml logs -f

down-tool: 
	@echo "Shutting down all anne-bonny containers"
	-@docker-compose -f docker-compose.tool.yml down

release:
	@echo "Firing new release helper"
	@npx release-it

test-report:
	@echo "Generating test report"
	@go test -coverprofile=coverage.out ./video-streamer/... ./storage/azure/...

test-review: test-report
	@go tool cover -html=coverage.out