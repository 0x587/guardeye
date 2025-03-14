FROM node:22.14.0-alpine

WORKDIR /build

COPY package*.json ./
RUN npm ci --omit=dev --registry=https://registry.npmmirror.com
RUN npm cache clean --force

COPY . .

EXPOSE 50051

CMD ["node", "index.js"]