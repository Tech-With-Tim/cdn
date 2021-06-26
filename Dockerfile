FROM golang:1.16-alpine3.13
WORKDIR /app
COPY . .
RUN go build -o main main.go

EXPOSE 5000
CMD /app/main migrate_up| true;/app/main runserver --host 0.0.0.0