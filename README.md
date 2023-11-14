# Odonto Dashboard

The Odonto Dashboard is a comprehensive solution designed for dental clinics, enabling efficient management of patient information and appointment details.

## Getting Started

To begin using the Odonto Dashboard, follow these steps:

1. **Clone the repository:**
   
```bash
git clone https://github.com/tomazcx/odonto-dashboard-go.git
```

## Running in Developer Mode

After cloning the repository, build the Docker image using the dedicated Dockerfile for the development environment:

```bash
docker build -t odonto-dashboard-dev -f Dockerfile.development .
```

Run the container:

```bash
docker run -p ${APP_PORT}:${APP_PORT} -v .:/usr/local/app --name odonto-dashboard odonto-dashboard-dev
```

Finally, initiate the bundle to enable autoreload for Tailwind CSS:

```bash
npm run  watch
```

## How to run it in production mode

Once the project is cloned locally, build the Docker image using the specified Dockerfile for the production environment:

```bash
docker build -t odonto-dashboard -f Dockerfile .
```

Run the container:

```bash
docker run -p ${APP_PORT}:${APP_PORT} -v .:/usr/local/app --name odonto-dashboard odonto-dashboard
```

This will set up the Odonto Dashboard in production mode, ready for efficient and secure use in a dental clinic setting.
