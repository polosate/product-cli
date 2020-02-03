FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

COPY . .

#RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o product-cli


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
ADD product.json /app/product.json
COPY --from=builder /app/product-cli .

CMD ["./product-cli"]
