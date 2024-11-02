# Imagen base
FROM node:14

WORKDIR /app
COPY . .

RUN npm install
RUN npm run build

EXPOSE 80
CMD ["npx", "serve", "-s", "build", "-l", "80"]
