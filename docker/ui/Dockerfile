# Étape de construction
FROM node:20-alpine AS builder

WORKDIR /app

# Copie des fichiers de dépendances
COPY package.json package-lock.json ./
RUN npm ci

# Copie du code source
COPY . .

# Construction de l'application
RUN npm run build

# Étape de production avec Nginx
FROM nginx:alpine

# Copie des fichiers de build
COPY --from=builder /app/dist /usr/share/nginx/html

# Copie de la configuration Nginx
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Exposer le port
EXPOSE 80

# Script pour remplacer la variable d'environnement API_URL au démarrage
RUN echo "#!/bin/sh" > /docker-entrypoint.d/30-env-replace.sh && \
    echo "sed -i \"s|__API_URL__|\${API_URL}|g\" /usr/share/nginx/html/assets/*.js" >> /docker-entrypoint.d/30-env-replace.sh && \
    chmod +x /docker-entrypoint.d/30-env-replace.sh

# Démarrage de Nginx
CMD ["nginx", "-g", "daemon off;"]
