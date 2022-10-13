COVER_MODULES = ./api/users/services/ ./api/tasks/services/ ./api/pokemons/services/

start:
	go run .

start.watch:
	air -d

start.prod:
	go build -o main && ./main

lint:
	@echo "FORMATTING"
	go fmt ./...
	@echo "LINTING: golangci-lint"
	golangci-lint run .

start.test:
	go test -v ./...

cover:
	@$(foreach dir,$(COVER_MODULES), \
		(cd $(dir) && \
		echo "[cover] $(dir)" && \
		go test  -coverprofile=cover.out -coverpkg=./... ./... && \
		go tool cover -html=cover.out -o cover.html) &&) true
