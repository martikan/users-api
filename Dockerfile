FROM golang:1.20-bookworm AS builder

WORKDIR /opt
COPY . .
RUN make build

FROM alpine:3.18 AS runner

WORKDIR /opt
COPY --from=builder /opt/users-api .

EXPOSE 3000
ENTRYPOINT ["/opt/users-api"]