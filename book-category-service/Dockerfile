FROM golang:1.22.5-alpine3.20 AS builder

LABEL author="Dwi Prasetiyo"
LABEL project="Synopsis Book Category Service"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./main ./cmd/main.go

# stage 2
FROM alpine:3.20  

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 1000 1100

CMD [ "./main" ]