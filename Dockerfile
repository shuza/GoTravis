FROM golang
COPY . /my-app
WORKDIR /my-app
RUN go mod vendor
ENTRYPOINT go run main.go
EXPOSE 8080
