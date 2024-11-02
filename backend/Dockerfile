# Imagen base
FROM golang:1.17

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o backend

EXPOSE 4000
CMD ["/app/backend"]
