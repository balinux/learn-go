
# Build stage
FROM golang:1.22-alpine AS build

# Install gcc and other build dependencies
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go application with optimizations
RUN CGO_ENABLED=1 go build -ldflags="-s -w" -o haircut-sqlite . && strip haircut-sqlite

# Final stage with alpine
FROM alpine

WORKDIR /

COPY --from=build /app/haircut-sqlite /haircut-sqlite
COPY --from=build /app/haircut.db /haircut.db

EXPOSE 8889

CMD ["/haircut-sqlite"]
