FROM node:20.9.0 as node_builder
WORKDIR /usr/local/node_app
COPY . .
RUN npm install
RUN npm run build

FROM golang:1.21.2 as go_builder
WORKDIR /usr/local/app
COPY . .
RUN go mod tidy
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN templ generate
COPY --from=node_builder /usr/local/node_app/public/dist/ /usr/local/app/public/dist
RUN go build -o ./temp/main
ENTRYPOINT [ "./temp/main" ]


