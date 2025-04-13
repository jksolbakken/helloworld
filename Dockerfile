FROM cgr.dev/chainguard/go AS builder
WORKDIR /app
COPY go.mod .
COPY ./cmd/main.go .
RUN CGO_ENABLED=0 go build -o hello main.go

FROM cgr.dev/chainguard/static
COPY --from=builder /app/hello /usr/bin/hello
EXPOSE 8080
ENTRYPOINT ["/usr/bin/hello"]
