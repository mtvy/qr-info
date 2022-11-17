
CONT_FLDR=deployments
BUILD_FILE=main
TESTS_FILE=test

COMMIT=save_changes

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

go-test:
	cd test && go build -o $(TESTS_FILE) . && cd ..
	test/test
	
enter:
	docker exec -it $(CONT) bash

git-push:
	git add .
	git commit -m "$(COMMIT)"
	git push
