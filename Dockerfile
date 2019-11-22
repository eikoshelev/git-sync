FROM golang:alpine AS build
WORKDIR /src
ADD . .
WORKDIR /src
RUN go build -o git-sync

FROM alpine
WORKDIR /bin
COPY --from=build /src/git-sync .
CMD ["git-sync"]
