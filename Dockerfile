FROM golang
COPY . /GoTravis
WORKDIR /GoTravis
RUN go mod vendor
ENTRYPOINT go run main.go
EXPOSE 8080
