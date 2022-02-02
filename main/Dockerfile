FROM node:17.3-alpine3.14
RUN mkdir -p app/public
COPY public/index.html app/public/
COPY package.json app/
COPY package-lock.json app/
COPY server.js app/
WORKDIR /app
RUN npm ci
EXPOSE 4000
CMD ["npm", "run", "start"]
