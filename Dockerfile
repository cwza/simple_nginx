# build stage
FROM golang:1.16-alpine AS build-env
ADD . /src
RUN cd /src && go mod download
RUN cd /src/producer && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o producer
RUN cd /src/consumer && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o consumer

# final stage
FROM alpine
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /src/producer/producer /
COPY --from=build-env /src/producer/producer.toml /
COPY --from=build-env /src/consumer/consumer /
COPY --from=build-env /src/consumer/consumer.toml /