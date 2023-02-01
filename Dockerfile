FROM golang:1.19-alpine

WORKDIR /app/

COPY ./ /app/

RUN go build -o ./service ./cmd/service

EXPOSE 8080

CMD [ "./service" ]