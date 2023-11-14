# Odonto Dashboard

A dashboard for a dentist clinic. Here you can manage your patients its appointments informations.

## How to run it in developer mode

First, clone the repository into your local machine:

```bash
    git clone https://github.com/tomazcx/odonto-dashboard-go.git
```

Then build the image with the designed Dockerfile for development environment:

```bash
    docker build -t odonto-dashboard-dev -f Dockerfile.development .
```

Run the container:

```bash
    docker run -p ${APP_PORT}:${APP_PORT} -v .:/usr/local/app --name odonto-dashboard odonto-dashboard-dev
```

Finally, run the bundle to work with autoreload for tailwindcss:

```bash
    npm run  watch
```

## How to run it in production mode

First, clone the repository into your local machine:

```bash
    git clone https://github.com/tomazcx/odonto-dashboard-go.git
```

Then build the image with the designed Dockerfile for production environment:

```bash
    docker build -t odonto-dashboard -f Dockerfile .
```

Run the container:

```bash
    docker run -p ${APP_PORT}:${APP_PORT} -v .:/usr/local/app --name odonto-dashboard odonto-dashboard
```
