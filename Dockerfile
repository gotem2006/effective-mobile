FROM golang:1.22-alpine as builder



WORKDIR /app

COPY go.mod . 

RUN go mod download

COPY . .


RUN go build cmd/app/main.go


FROM alpine:3


COPY --from=builder /app/main .
COPY --from=builder /app/config.env .
COPY --from=builder /app/pkg/database/migrations ./pkg/database/migrations

ENTRYPOINT ["./main"]