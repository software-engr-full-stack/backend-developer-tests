name := stackpath-rest-service
path := people
source_file := $(abspath ./tmp/${name})
output_path := ${source_file}.zip

_args := -var="name=${name}" -var="source_file=${source_file}" -var="output_path=${output_path}" -var="path=${path}"

.PHONY: build-service
build-service:
	mkdir -p $(dir ${source_file}) && \
	cd ../../../cmd/microservice/app/handler && \
	GOOS=linux GOARCH=amd64 go build -o "${source_file}"

.PHONY: apply-service
apply-service: build-service
	terraform init && \
	terraform plan  ${_args} && \
	terraform apply ${_args} -auto-approve

.PHONY: destroy-service
destroy-service:
	terraform destroy ${_args} -auto-approve

test-service:
	curl "$$(terraform output -raw base_url)/"

test-service-id:
	# TODO: not working
	curl "$$(terraform output -raw base_url)/df12ce76-767b-4bf0-bccb-816745df9e70"

test-service-names:
	curl "$$(terraform output -raw base_url)/?first_name=John&last_name=Doe"

test-service-phone:
	curl "$$(terraform output -raw base_url)/?phone_number=%2b1+(800)+555-1212"

test: test-service test-service-names test-service-phone
