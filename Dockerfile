FROM node:alpine as docs

WORKDIR /docs

COPY ./docs/docs-template/package.json /docs
COPY ./docs/docs-template/package-lock.json /docs

RUN npm install

COPY ./docs/docs-template /docs

RUN npm run build

# ---------

FROM golang:1.16-alpine3.13 as builder
WORKDIR /app

ADD go.mod .
ADD go.sum .
RUN go mod download -x

COPY . .
COPY --from=docs /docs/public/ /app/docs/docs-template/public/

RUN go run main.go generate_docs
RUN CGO_ENABLED=0 GOOS=linux go build -o main -ldflags "-w -s"

# --------

FROM alpine
WORKDIR /app
COPY --from=builder /app/ /app/
RUN chmod 755 ./main

EXPOSE 5000
CMD ["/app/main", "runserver", "--host", "0.0.0.0"]
