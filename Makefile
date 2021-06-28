# constants
NAME=app
PROJECT?=github.com/MarkoLuna/GoEmployeeCrud

build:
	go build -o ${NAME} "${PROJECT}/cmd/server"

test:
	go test -race "${PROJECT}/..."

test-cover:
	go test -cover "${PROJECT}/..."

test-total-cover:
	go test "${PROJECT}/..." -coverprofile cover.out > /dev/null
	go tool cover -func cover.out
	rm cover.out

run: build
	./app

clean:
	go clean "${PROJECT}/..."
	rm -f ${NAME}

docker-build:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${NAME} "${PROJECT}/cmd/server"
	docker build -t goemployee_crud:latest .
	rm -f ${NAME}

docker-run: docker-build
	docker run -it -p 8080:8080 --rm goemployee_crud

docker-compose-run: docker-build
	docker-compose up

