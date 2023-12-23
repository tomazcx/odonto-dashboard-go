FROM node:20.9.0 as node_builder
WORKDIR /usr/local/node_app
COPY . .
RUN npm install
RUN npm run build

FROM golang:1.21 as go_builder
WORKDIR /usr/local/app
COPY . .
RUN go install github.com/a-h/templ/cmd/templ@latest
COPY --from=node_builder /usr/local/node_app/public/dist/ /usr/local/app/public/dist
ENTRYPOINT [ ".docker/entrypoint.sh" ]


