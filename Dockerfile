FROM golang:latest AS builder

# https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
RUN apt-get update; \
    apt-get install -y --no-install-recommends \
    git;\
    rm -rf /var/lib/apt/lists/*

COPY . /app

WORKDIR /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o \
    /go/bin/release /app/cmd/api

# https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/static

# change this feature using go:embed
COPY --from=builder /app/.env /app/.env
COPY --from=builder /go/bin/release /go/bin/release

ENV DATABASE_URL=""
ENV DISCORD_TOKEN=""

EXPOSE 8080
ENTRYPOINT ["/go/bin/release"]