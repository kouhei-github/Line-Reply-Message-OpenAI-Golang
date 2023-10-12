FROM golang:1.20 as build

WORKDIR /app

RUN go install github.com/cosmtrek/air@v1.29.0

COPY ./src/go.mod /app
COPY ./src/go.sum /app
COPY ./src /app
COPY ./src/zeus-firebase-admin.json /app/zeus-firebase-admin.jsonq
RUN CGO_ENABLED=0 go build -tags lambda.norpc -o main main.go

