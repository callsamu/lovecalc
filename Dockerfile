FROM golang:1.17-alpine AS builder
RUN mkdir /app

ADD . /build
WORKDIR /build

RUN go mod download
RUN go build -o main cmd/web/*

FROM alpine AS runner
COPY --from=builder /build/main /bin/main
ENTRYPOINT ["/bin/main", "port=:4000"]
