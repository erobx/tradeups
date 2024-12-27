FROM node:23-alpine

WORKDIR /app
RUN npm install -g pnpm
COPY . .
RUN pnpm install

CMD ["pnpm", "run", "dev", "--host"]
