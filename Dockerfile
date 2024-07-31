FROM golang

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/app

RUN go build -o /app/user-service .

CMD ["/app/user-service"]


