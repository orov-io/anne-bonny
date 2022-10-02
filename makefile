up:
	@echo "Booting anne-bonny services with docker-compose"
	@docker-compose up -d --build

logs:
	@echo "Tailing anne-bonny containers logs"
	@docker-compose logs -f

down: 
	@echo "Shutting down all anne-bonny containers"
	-@docker-compose down

reboot: down up logs