# Odonto Dashboard

The Odonto Dashboard is a comprehensive solution designed for dental clinics, enabling efficient management of patient information and appointment details.

## Getting Started

To begin using the Odonto Dashboard, start by cloning the repository:
   
```bash
git clone https://github.com/tomazcx/odonto-dashboard-go.git
```

## Running in Developer Mode

After cloning the repository, create an `.env` file based on `.env.example` and assign values to each variable according to its purpose.

Once all the variables are correctly assigned, initiate the development server using `docker-compose`:

```bash
docker-compose up
```

Finally, in another terminal instance, initiate the bundle to enable autoreload for Tailwind CSS:

```bash
npm run watch
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
Don't forget to assign the correct values for `.env` variables, specifying the production database credentials, host and name.
