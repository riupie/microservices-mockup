FROM node:16 as builder

WORKDIR /usr/src/app
COPY . .

RUN npm install && npm run coverage

FROM gcr.io/distroless/nodejs:16

WORKDIR /app

COPY --from=builder /usr/src/app /app

EXPOSE 3001
CMD ["server.js"]
