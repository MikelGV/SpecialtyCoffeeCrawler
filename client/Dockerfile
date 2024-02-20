FROM node:latest

WORKDIR ./app

RUN npm install typescript

COPY ./package.json ./

RUN npm install

COPY . .

EXPOSE 3000

