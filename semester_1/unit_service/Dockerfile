FROM golang:alpine

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /server ./cmd/main.go

CMD ["/server"]
