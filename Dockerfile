FROM golang:1.19

WORKDIR /app

COPY . /app

EXPOSE 8000

CMD [ "go", "run","/app/src/main.go" ]
