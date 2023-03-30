FROM node:latest

WORKDIR /frontend

COPY . /frontend/

ENV PORT 3000

RUN yarn

EXPOSE ${PORT}

RUN chmod +777 start.sh

CMD ["./start.sh"]