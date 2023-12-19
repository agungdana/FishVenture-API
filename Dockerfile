FROM golang:1.19.11-alpine3.18 AS builder
WORKDIR /builder
COPY . .
RUN go build -o ./gobin

FROM alpine:3.16 AS executable
WORKDIR /App
EXPOSE 8080
ARG APP_ENV
COPY --from=builder /builder/gobin .
COPY --from=builder /builder/.env .
COPY --from=builder /builder/flutter-firebase-79d6e-firebase-adminsdk-t143e-04bf70fe9b.json .
ENTRYPOINT ["./gobin"]
