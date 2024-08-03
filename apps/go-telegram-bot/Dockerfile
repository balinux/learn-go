FROM docker.io/library/golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o telegram-go .

# add stage for scratch image
FROM scratch
COPY --from=build /app/telegram-go /telegram-go
CMD [ "/telegram-go" ]
