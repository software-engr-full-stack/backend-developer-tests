# About

Answers to stackpath.com's Backend Developer Tests

Original [README](https://github.com/stackpath/backend-developer-tests)

# Answers

## Unit Testing

The answer is just one test file located [here](https://github.com/software-engr-full-stack/backend-developer-tests/blob/master/unit-testing/pkg/fizzbuzz/fizzbuzz_test.go).

## Web Services

The core answer is in the [people.go](https://github.com/software-engr-full-stack/backend-developer-tests/blob/master/rest-service/pkg/handlers/people.go) file inside the `rest-api/pkg/handlers` sub-directory. But I've thrown a bunch of extras so the answer to this test is pretty much the whole [rest-service](https://github.com/software-engr-full-stack/backend-developer-tests/tree/master/rest-service) sub-directory. Extras:

1. Lots of tests for the main handler in the `rest-api/pkg/handlers` sub-directory.

2. The service is deployed using AWS Lambda. The endpoint is [here](https://xflrazp4w3.execute-api.us-west-1.amazonaws.com/people/).
`cd` into `rest-service/cmd/microservice/deploy`, then run `make test` to test the different scenarios (GET /people, GET /people/?first_name=...&last_name=...) specified in the test against the endpoint. I used Terraform to provision the AWS resources. The Terraform code is located [here](https://github.com/software-engr-full-stack/backend-developer-tests/tree/master/rest-service/cmd/microservice/deploy).

3. Instrumentation and metrics gathering using Prometheus and Grafana. Some of the metrics gathered are total number of requests and response time. Use the term `stackpath_dev_tests` in Grafana's Metrics Browser to filter the custom metrics I've added. For now, this only works locally.

#### Grafana

1. `cd` into `rest-service`, then run `docker compose up`. This command will run the REST service, Prometheus, Grafana, and Node Exporter inside containers. Once the containers are running, proceed to the next steps.

2. Open your browser to [http://localhost:3030/](http://localhost:3030/). This is the Grafana portal. The default log in username and password is `admin`.

3. Once you're logged in, click on "Add your first data source", then click on "Prometheus", then enter "http://prometheus:9090" in the HTTP URL text box. Click on "Save & Test" button located below. If everything works, you should see a green check mark "Data source is working" status message.

4. Click on "Explore" on the side bar (compass icon), click on "Metrics Browser", then enter "stackpath_dev_tests" in the "Select a metric" text box. Click on a metric to see information and graphs about that metric.

## Input Processing

The [answer](https://github.com/software-engr-full-stack/backend-developer-tests/blob/master/input-processing/pkg/inputprocessor/inputprocessor.go) and the [test file](https://github.com/software-engr-full-stack/backend-developer-tests/blob/master/input-processing/test/pkg/inputprocessor/inputprocessor_test.go).

## Concurrency

[Here is the answer to the simple pool test.](https://github.com/software-engr-full-stack/backend-developer-tests/blob/master/concurrency/pkg/concurrency/simple_pool.go)
