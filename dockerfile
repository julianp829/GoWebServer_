# Use a lightweight Go image as the base
FROM golang:1.23 AS builder

# Set the working directory in the container
WORKDIR /build

# Copy all files into build directory
COPY . .


RUN go mod download
RUN go build -o ./gws

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /build/gws ./gws
CMD ["/app/gws"]
