# Builder
FROM golang:1.22-alpine3.19 AS builder
ARG VERSION

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
WORKDIR /migrate

RUN cp "$(go env GOPATH)/bin/migrate" migrate

# Runner
FROM alpine:3.19

RUN apk add --no-cache ca-certificates

COPY --from=builder /migrate/migrate /usr/local/bin/migrate
RUN ln -s /usr/local/bin/migrate /migrate

ENTRYPOINT ["migrate"]
CMD ["--help"]