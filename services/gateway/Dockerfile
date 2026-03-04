FROM golang:1.26-alpine AS builder

WORKDIR /build

COPY go.mod go.sum* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /gateway ./cmd/...

FROM scratch
COPY --from=builder /gateway /gateway
ENTRYPOINT ["/gateway"]
