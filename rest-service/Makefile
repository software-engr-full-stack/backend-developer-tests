.PHONY: test
test:
	go test -v -cover ./... && \
	golangci-lint run

build_path := /tmp/stackpath.io-backend-developer-tests-rest-service
.PHONY: build-and-run
build-and-run:
	go build -o ${build_path} && ${build_path}

.PHONY: live-reload
live-reload:
	npx nodemon --signal SIGTERM --ext go --exec 'make build-and-run' && \
		notify-send --expire-time 12000 '... live reload exited'
