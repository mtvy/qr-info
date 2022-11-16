
CONT_FLDR=deployments
BUILD_FILE=main


info:
	docker ps -a

up:
	cd $(CONT_FLDR) && docker-compose up -d

build:
	cd $(CONT_FLDR) && docker-compose build

down:
	cd $(CONT_FLDR) && docker-compose down

go-push: go-build
	cmd/main

go-run:
	cmd/main

go-build:
	cd cmd && go build -o $(BUILD_FILE) . && cd ..
	
enter:
	docker exec -it $(CONT) bash
