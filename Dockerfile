FROM golang

WORKDIR /app

COPY . .

WORKDIR /app/cmd

RUN go build -o /app/user-service .

CMD ["/app/user-service"]


