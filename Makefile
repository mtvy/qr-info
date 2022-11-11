
info:
	docker ps -a

up:
	cd deployments && docker-compose --project-name="qr-info" up -d

down:
	cd deployments && docker-compose --project-name="qr-info" down

enter:
	docker exec -it $(CONT) bash
