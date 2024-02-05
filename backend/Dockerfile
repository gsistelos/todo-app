FROM golang:latest

WORKDIR /app

COPY src .

RUN go build -o bin/todo-app

# Script to wait for a host:port to be available
RUN curl -fsSLO https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh
RUN chmod +x wait-for-it.sh

ENTRYPOINT ["./wait-for-it.sh", "mysql:3306", "--", "./bin/todo-app"]
