FROM golang:1.25.1-alpine3.22 AS builder

RUN apk update && \
  apk upgrade --no-cache && \
  apk add --no-cache \
  git=2.45.2-r0 \
  ca-certificates=20240705-r0 \
  tzdata=2024b-r0 && \
  rm -rf /var/cache/apk/*

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && \
  go mod verify

COPY . .

ARG VERSION=dev
ARG COMMIT_SHA=unknown
ARG BUILD_DATE
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
  -a -installsuffix cgo \
  -ldflags="-w -s -X main.version=${VERSION} -X main.commitSHA=${COMMIT_SHA} -X main.buildDate=${BUILD_DATE}" \
  -o andino ./cmd/server

FROM alpine:3.22

RUN apk --no-cache add \
  ca-certificates=20240705-r0 \
  tzdata=2024b-r0 \
  wget=1.24.5-r0

RUN addgroup -S andino && adduser -S andino -G andino

WORKDIR /app

COPY --from=builder /app/andino .
COPY --from=builder /app/configs ./configs

RUN chown -R andino:andino /app

USER andino

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health 2>/dev/null || exit 1

CMD ["./andino"]
