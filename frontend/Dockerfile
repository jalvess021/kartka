# Base para desenvolvimento
FROM node:20.15.1 AS dev

WORKDIR /vue/app

COPY ./package*.json ./

RUN npm install

COPY . .

CMD ["npm", "run", "dev"]

# Base para produção
FROM node:20.15.1 AS build

WORKDIR /vue/app

COPY ./package*.json ./

RUN npm install
COPY . .

RUN npm run build

# Servir o build estático usando NGINX
FROM nginx:alpine AS production

COPY --from=build /vue/app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
