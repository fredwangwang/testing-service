FROM golang:buster AS build

WORKDIR /app

COPY . .

RUN go build -o testing-service

FROM ubuntu

RUN apt update && apt install -y curl

COPY --from=build /app/testing-service /usr/local/bin/testing-service

CMD [ "/usr/local/bin/testing-service" ]
