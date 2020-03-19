FROM golang:alpine AS builder

RUN mkdir /app

ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/main main.go ./...

FROM alpine:latest AS production
# Copy contents from the builder stage
COPY --from=builder /app .

CMD [ "./cmd/main" ]
