FROM golang:1.14 AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/foobar


FROM alpine AS runner

COPY --from=builder /bin/foobar /bin/foobar
CMD ["/bin/foobar"]
