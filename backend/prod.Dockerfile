# FROM golang:1.22.5 as builder
# # Set the working directory inside the container
# WORKDIR /app

# # Copy go.mod and go.sum files
# COPY go.mod go.sum ./

# # Download dependencies
# RUN go mod download

# # Copy the rest of the application code
# COPY . .

# # Build the Go application
# RUN CGO_ENABLED=1 GOOS=linux go build -o main .



# ENV GIN_MODE=release

# CMD ["./main"]



# # Stage 2: Create a minimal image for the application
# FROM alpine:3.20
# RUN apk add --no-cache libc6-compat


# # Copy the Go binary from the builder stage
# COPY --from=builder /app/main /main

# ENV GIN_MODE=release

# CMD ["/main"]



FROM golang:1.22.5-alpine3.20 as builder
ENV CGO_ENABLED=1
WORKDIR /app
COPY . .
RUN apk add --no-cache --update git build-base
RUN go mod tidy \
    && go build -o main .


FROM alpine:latest as runner
ENV TZ=Asia/Tokyo
RUN apk update
RUN apk --no-cache add ca-certificates tzdata libc6-compat libgcc libstdc++

COPY --from=builder /app/main /app/main

ENV GIN_MODE=release

CMD ["/app/main"]