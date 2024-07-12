FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest AS runtime

WORKDIR /app

COPY --from=build /app/main .

EXPOSE 3000

CMD ["/app/main"]