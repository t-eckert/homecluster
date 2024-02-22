FROM golang:1.20.3-alpine AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN GOOS=linux GOARCH=arm64 go build -o main .

ENV OPEN_WEATHER_API_KEY=
ENV LONGITUDE=
ENV LATITUDE=

EXPOSE 8080
CMD ["/app/main"]
