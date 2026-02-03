FROM golang:1.24 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o server ./server && chmod +x server

FROM alpine:latest
COPY --from=build --chmod=755 /app/server /server
EXPOSE 8080
CMD ["/server"]
