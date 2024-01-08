# golang build image
FROM golang:latest AS go

WORKDIR /app
COPY ./ /app/
RUN go mod download && CGO_ENABLED=0 go build -o main ./cmd/main.go

# golang runtime image
FROM alpine:latest AS runtime
WORKDIR /app
COPY --from=go /app/main /app/main
USER 1001
EXPOSE 8080
RUN readlink -f main
CMD [ "/app/main" ]

# docker build -t chat-app-server:0.0.1 .
# docker run -d --rm --name chat-app-server -p 8080:8080 chat-app-server
