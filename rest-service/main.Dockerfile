FROM golang:1.18-bullseye as builder

ENV APP_HOME /app

WORKDIR "$APP_HOME"
COPY ./ .

RUN go mod download
RUN go mod verify
RUN go build -o stackpath-backend-developer-tests-rest-service

FROM golang:1.18-bullseye

ENV APP_HOME /app
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

# RUN go build -o stackpath-backend-developer-tests-rest-service
COPY --from=builder "$APP_HOME"/stackpath-backend-developer-tests-rest-service $APP_HOME

EXPOSE 8000
CMD ["/app/stackpath-backend-developer-tests-rest-service"]
