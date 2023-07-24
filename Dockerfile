# Build stage
FROM golang:1.17-alpine AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Deploy stage 
FROM alpine:3.14  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /app/main .
COPY --from=build /app/app.env .
CMD ["./main"]

EXPOSE 8080