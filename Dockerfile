FROM golang:latest AS build

WORKDIR /build
COPY . .

ENV CGO_ENABLED=0
RUN go mod download
RUN go build -o bin/app cmd/web/

FROM scratch AS bin
COPY --from=build /build/bin/app /app
ENTRYPOINT ["/app"]
