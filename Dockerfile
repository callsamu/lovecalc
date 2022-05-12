FROM golang:latest AS base
WORKDIR /build
COPY . .
RUN go mod download
ENV CGO_ENABLED=0

FROM base AS build
RUN go build -o bin/app /build/cmd/web/

FROM base AS test
RUN go test -v ./...

FROM scratch AS bin
COPY --from=build /build/bin/app /app
ENTRYPOINT ["/app"]
