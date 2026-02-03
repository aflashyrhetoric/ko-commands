FROM golang:1.24 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/server ./server

FROM alpine:latest
COPY --from=build /bin/server /bin/server
EXPOSE 8080
CMD ["/bin/server"]
