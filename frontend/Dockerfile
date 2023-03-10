FROM node:latest

WORKDIR /frontend

COPY . /frontend/

RUN yarn

RUN yarn build

ENV PORT 3000

EXPOSE ${PORT}

CMD ["yarn", "start"]