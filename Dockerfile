FROM node:20.9.0 as node_builder
WORKDIR /usr/local/node_app
COPY . .
RUN npm install
RUN npm run build

FROM golang:1.21.2 as go_builder
WORKDIR /usr/local/app
COPY . .
RUN go mod tidy
COPY --from=node_builder /usr/local/node_app/dist/ /usr/local/app/dist
RUN go build -o ./temp/main
ENTRYPOINT [ "./temp/main" ]


