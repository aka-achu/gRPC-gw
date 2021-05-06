FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o gw ./main.go

FROM scratch
WORKDIR /app/
COPY --from=builder /app/gw  .
EXPOSE 8080
EXPOSE 8081
CMD ["./gw"]