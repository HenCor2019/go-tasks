COVER_MODULES = ./api/tasks/services/ ./api/users/services/

start:
	go run .

start.watch:
	air -d

start.prod:
	go build -o main && ./main

lint:
	go fmt . && golangci-lint run .

start.test:
	go test -v ./...

cover:
	@$(foreach dir,$(COVER_MODULES), \
		(cd $(dir) && \
		echo "[cover] $(dir)" && \
		go test  -coverprofile=cover.out -coverpkg=./... ./... && \
		go tool cover -html=cover.out -o cover.html) &&) true
