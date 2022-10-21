ifneq ("$(wildcard $(local.env))","")
    include local.env
	export $(shell sed 's/=.*//' local.env)
endif

ifneq ("$(wildcard $(override.env))","")
    include override.env
	export $(shell sed 's/=.*//' override.env)
endif

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