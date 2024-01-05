FROM golang:1.21.5-bullseye AS build

RUN apt-get update && apt-get install -y git

WORKDIR /app

RUN echo product-service

RUN echo product-service

RUN git clone https://github.com/akshay0074700747/product-service-grpc.git .

RUN go mod download

WORKDIR /app/cmd

RUN go build -o bin/product-service

COPY /cmd/.env /app/cmd/bin/

FROM busybox:latest

WORKDIR /product-service

COPY --from=build /app/cmd/bin/product-service .

COPY --from=build /app/cmd/bin/.env .

EXPOSE 50004

CMD ["./product-service"]