FROM golang:latest

COPY main.go /app/main.go

CMD ["go", "run", "/app/main.go"]