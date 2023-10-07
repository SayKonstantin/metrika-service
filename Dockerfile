FROM golang:alpine AS builder
WORKDIR /build
COPY . .
RUN go build -o metrika /build/cmd/server/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /build/metrika ./
RUN chmod +x metrika
RUN apk add --no-cache tzdata
ENV TZ=Europe/Moscow
ENTRYPOINT ["./metrika", "--env"]