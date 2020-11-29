FROM golang:1.15.5-buster AS dev

WORKDIR /code
RUN go get github.com/githubnemo/CompileDaemon && \
    go get -u golang.org/x/lint/golint

COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN GOOS=linux go build -a -o /app .

CMD CompileDaemon -log-prefix=false -build="go build -a -o /app" -command="/app"

FROM scratch AS prod
COPY --from=dev /app /app
ENTRYPOINT ["/app"]
