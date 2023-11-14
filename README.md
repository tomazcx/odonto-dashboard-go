# Odonto Dashboard

A dashboard for a dentist clinic. Here you can manage your patients its appointments informations.

To start using it, clone the repository:

```bash
git clone https://github.com/tomazcx/odonto-dashboard-go.git
```

## How to run it in developer mode

After cloning the repository, build the image with the designed Dockerfile for development environment:

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

After cloning the project locally, build the image with the designed Dockerfile for production environment:

```bash
docker build -t odonto-dashboard -f Dockerfile .
```

Run the container:

```bash
docker run -p ${APP_PORT}:${APP_PORT} -v .:/usr/local/app --name odonto-dashboard odonto-dashboard
```
