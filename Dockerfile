FROM golang:1.14.3-alpine

ENV CGO_ENABLED 0

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

# download all dependencies
RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]