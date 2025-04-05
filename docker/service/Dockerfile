FROM golang:1.21-alpine AS builder

WORKDIR /app

# Installation des dépendances système
RUN apk add --no-cache git ca-certificates tzdata

# Copie des fichiers Go
COPY go.mod go.sum ./
RUN go mod download

# Copie du code source
COPY . .

# Compilation de l'application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o navidrome-mix ./cmd/service/main.go

# Image finale
FROM alpine:3.19

WORKDIR /app

# Installation des dépendances nécessaires
RUN apk add --no-cache ca-certificates tzdata

# Copie du binaire compilé depuis le builder
COPY --from=builder /app/navidrome-mix .

# Copie des fichiers de configuration
COPY --from=builder /app/config ./config

# Exposer le port de l'API
EXPOSE 8085

# Définir le point d'entrée
ENTRYPOINT ["/app/navidrome-mix"]
