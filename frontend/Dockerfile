# Multi-stage build for production
FROM node:20-alpine as builder

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm ci

COPY . .

RUN npm run build

# Production stage
FROM node:20-alpine

WORKDIR /app

COPY --from=builder /app/dist ./dist
COPY --from=builder /app/package.json ./

RUN npm ci --production

EXPOSE 3000

CMD ["npm", "run", "preview"]
