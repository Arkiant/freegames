FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo /app/cmd/freegames

FROM alpine:latest 

RUN apk --no-cache add ca-certificates

WORKDIR /

ENV DATABASE_URL=""
ENV DISCORD_TOKEN=""

COPY --from=builder /app/freegames .

CMD ["./freegames"]