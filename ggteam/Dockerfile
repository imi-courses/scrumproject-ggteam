FROM node:20-alpine

WORKDIR /frontend

EXPOSE 8080

COPY package.json package-lock.json ./

RUN npm i --silent

COPY . ./

CMD ["npm", "run", "docker"]
