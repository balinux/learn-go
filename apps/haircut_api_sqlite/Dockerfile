FROM docker.io/library/golang:1.22-alpine AS build
# Install gcc and other build dependencies
RUN apk add. --no-cache gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o haircut-sqlite .

# add stage for scratch image
FROM scratch
COPY --from=build /app/haircut-sqlite /haircut-sqlite
COPY --from=build /app/haircut.db /haircut.db
EXPOSE 8889
CMD [ "/haircut-sqlite" ]
