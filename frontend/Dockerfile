FROM node:15.3.0-alpine3.12

WORKDIR /code
COPY ./package.json ./yarn.lock ./

RUN yarn install
COPY \
  ./babel.config.js \
  ./package.json \
  ./public \
  ./src \
  ./yarn.lock \
  ./


CMD yarn run serve

