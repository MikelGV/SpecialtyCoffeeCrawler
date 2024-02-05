FROM node:alpine

WORKDIR /client

RUN npm install typescript

COPY .package.json ./
COPY .package-lock.json ./

RUN npm install

COPY . .

CMD ["npm", "start"]
