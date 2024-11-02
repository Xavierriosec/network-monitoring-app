# Imagen base
FROM golang:1.19

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o backend

EXPOSE 4000
CMD ["/app/backend"]
