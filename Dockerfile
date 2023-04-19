FROM golang:1.19 AS builder
WORKDIR /src
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -o git-sync ./cmd/git-sync/

FROM alpine:latest AS production
WORKDIR /bin
COPY --from=builder /src/git-sync .
CMD ["git-sync"]