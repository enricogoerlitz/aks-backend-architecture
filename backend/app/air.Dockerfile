FROM golang:1.23.2-alpine

WORKDIR /app

RUN apk add --no-cache curl git && \
    go install github.com/air-verse/air@latest

COPY . /app/
RUN go mod download

# Run Air in development mode
CMD ["air"]