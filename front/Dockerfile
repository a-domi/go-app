FROM node:20-alpine

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install
RUN npm install axios
RUN apk add bash

COPY . .

CMD [ "npm", "start" ]