_port := 8100

# AWS Lambda running locally for testing
.PHONY: run-local
run-local:
	sam build && sam local start-api --port ${_port}

.PHONY: test-local
test-local:
	curl http://localhost:${_port}/people

.PHONY: run-local test
