FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o idonce-web .
RUN ./idonce-web --build

FROM caddy:2-alpine
COPY --from=builder /app/dist /srv
COPY Caddyfile.static /etc/caddy/Caddyfile
EXPOSE 9091
