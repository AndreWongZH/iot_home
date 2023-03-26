FROM golang:latest

WORKDIR /backend

RUN apt-get update
RUN apt-get install nmap -y

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/main .

ENV PORT=3001

EXPOSE ${PORT}

CMD ["main"]