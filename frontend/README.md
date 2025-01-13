This is a Band manager project created with Next.js (https://nextjs.org).

## Getting Started

First, configure the .env file:

```bash
NEXTAUTH_SECRET={your_hash}
NEXTAUTH_URL=http://localhost:{frontend_port}/
NEXT_PUBLIC_BACKEND_HOST={backend_host}
NEXT_PUBLIC_BACKEND_PORT={backend_port}

```
To run this code in development mode build docker container with development.Dockerfile. If you would like to compile the code for production usage use production.Dockerfile.
