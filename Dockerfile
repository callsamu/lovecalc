# syntax = docker/dockerfile:1-experimental

FROM golang:latest AS base
WORKDIR /build
COPY . .
RUN go mod download
ENV CGO_ENABLED=0

FROM base AS test
ENTRYPOINT ["go", "test", "-v", "./..."]

FROM base AS build
RUN --mount=type=cache,target=/root/.cache/go-build \ 
go build -o bin/app /build/cmd/web/

FROM scratch AS bin
COPY --from=build /build/bin/app /app
COPY --from=build /build/ui /ui
ENTRYPOINT ["/app"]
