FROM golang as build
WORKDIR /httpserver/
COPY . .
ENV GO111MODULE=off \
	CGO_ENABLED=0 \
        GOOS=linux \
        GOARCH=amd64
RUN go build -o server

FROM alpine
COPY --from=build httpserver/server httpserver
ENTRYPOINT ["./httpserver"]
