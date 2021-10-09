FROM golang:1.16-alpine as builder
RUN apk --no-cache add ca-certificates git
WORKDIR /build


# Fetch dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . ./
RUN CGO_ENABLED=0 go build api/main.go

# Create final image
FROM alpine
ENV DB_URL=<DB_URL>
WORKDIR /
COPY --from=builder /build/main .
EXPOSE 8080
CMD ["./main"]