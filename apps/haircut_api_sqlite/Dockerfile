FROM docker.io/library/golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o haircut-sqlite .

# add stage for scratch image
FROM scratch
COPY --from=build /app/haircut-sqlite /haircut-sqlite
COPY --from=build /app/haircut.db /haircut.db
EXPOSE 8889
CMD [ "/haircut-app" ]
