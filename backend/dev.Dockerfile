# Use the official Golang image as a builder
FROM public.ecr.aws/docker/library/golang:1.22 as builder

RUN mkdir /app
# Set the working directory inside the container
WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]

