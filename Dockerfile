FROM golang:alpine AS builder

WORKDIR /service
COPY . .
RUN go mod vendor
RUN go build


FROM alpine

WORKDIR /service
COPY --from=builder /service /service
CMD ["/service/go-mysql-template"]