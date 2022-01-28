# Step 1: Modules caching
FROM golang:1.17.1-alpine as modules
WORKDIR /modules
COPY go.mod go.sum /
RUN go mod download

# Step 2: Builder
FROM golang:1.17.1-alpine as builder
COPY --from=modules /go/pkg /go/pkg
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-extldflags=-static" -o /bin/app ./cmd/server

# Step 3: Run
FROM scratch
COPY --from=builder /bin/app /app
COPY --from=builder /app/migrations /migrations
CMD ["/app"]