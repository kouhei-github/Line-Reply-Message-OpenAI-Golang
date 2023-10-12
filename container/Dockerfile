FROM golang:1.20 as build

WORKDIR /app

RUN go install github.com/cosmtrek/air@v1.29.0

COPY ./src/go.mod /app
COPY ./src/go.sum /app
COPY ./src /app
RUN go build -tags lambda.norpc -o main main.go


#FROM public.ecr.aws/lambda/provided:al2
#COPY --from=build /app/main ./main
#ENTRYPOINT [ "./main" ]
