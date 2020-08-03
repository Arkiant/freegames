FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest 

RUN apk --no-cache add ca-certificates

WORKDIR /

ENV DATABASE_URL=mongodb://localhost:27017

COPY --from=builder /app/app .

CMD ["./app"]