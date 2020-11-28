FROM golang:1.15 AS dev

WORKDIR /code
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN GOOS=linux go build -a -o /app .

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -a -o /app" -command="/app"

FROM scratch
COPY --from=dev /app /app
ENTRYPOINT ["/app"]
