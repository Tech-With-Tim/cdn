FROM golang:1.16-alpine3.13 as builder
WORKDIR /app

ADD go.mod .
ADD go.sum .
RUN go mod download -x

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main -ldflags "-w -s"

FROM alpine
WORKDIR /app
COPY --from=builder /app/main /app/main
RUN chmod 755 ./main

EXPOSE 5000
CMD ["/app/main", "runserver", "--host", "0.0.0.0"]