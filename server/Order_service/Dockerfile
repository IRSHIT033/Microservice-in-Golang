FROM golang:alpine

RUN apk update && apk add gcc g++ libc-dev librdkafka-dev pkgconf

# Move to working directory /build
WORKDIR /server/Order_service

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -tags musl -o main ./cmd/


# Export necessary port
EXPOSE 3002

# Command to run when starting the container
CMD ["/server/Order_service/main"]