FROM node:20.9.0 as node_builder
WORKDIR /usr/local/node_app
COPY . .
RUN npm install
RUN npx webpack build

FROM golang:1.21.2 as go_builder
WORKDIR /usr/local/app
COPY . .
RUN go mod tidy
COPY --from=node_builder /usr/local/node_app/dist/ /usr/local/app/dist
ENTRYPOINT [ "go" ]
CMD ["run", "."]


