# Band Manager

Band Manager is a web application for music band management, built with Next.js frontend, Go backend, and PostgreSQL database.

## Technologies

- **Frontend**: Next.js
- **Backend**: Go
- **Database**: PostgreSQL 16
- **Containerization**: Docker
- **SMTP**: Email notifications via Gmail

## Requirements

- Docker
- Docker Compose
- Node.js (for frontend development)
- Go (for backend development)

## Installation and Setup

1. Clone the repository:

```bash
git clone https://github.com/PZSP2-Team/band-manager.git
cd band-manager
```

2. Copy `.env.example` to `.env` and fill in the required environment variables:

```bash
cp .env.example .env
```

3. Launch the application using Docker Compose:

```bash
docker-compose up -d --build
```

The application will be available at:

- Frontend: `http://localhost:3000`

Backend is only available through frontend, because next.js uses proxy pass to backend requests and JWT verification

## Environment Variables

### Database

- `POSTGRES_DB` - database name
- `POSTGRES_USER` - database user
- `POSTGRES_PASSWORD` - database password
- `POSTGRES_PORT` - PostgreSQL port (default: 5432)

### Backend

- `BACKEND_PORT` - backend server port (default: 8080)
- `BACKEND_HOST` - backend server host

### Frontend

- `FRONTEND_PORT` - frontend server port (default: 3000)
- `FRONTEND_HOST` - frontend server host
- `NODE_ENV` - Node.js environment (development/production)

### Email

- `SMTP_HOST` - SMTP server host (smtp.gmail.com)
- `SMTP_PORT` - SMTP port (587)
- `EMAIL_FROM` - notification sender email address
- `APP_PASSWORD` - Gmail application password

## Project Structure

```
band-manager/
├── frontend/           # Next.js application
├── backend/           # Go server
├── docker-compose.yml # Docker Compose configuration
└── .env              # Environment variables
```

## Security

- Application uses internal Docker network for service communication
- PostgreSQL data is stored in a separate Docker volume
- Database healthchecks implementation

## Data Persistence

The application uses two Docker volumes:

- `postgres_data`: for PostgreSQL database data
- `notesheet_files`: for uploaded application files

## Authors

- Mikołaj Rożek (@mikorozek) - Full Stack Developer
- Marcin Lisowski (@mnlisows) - Backend Developer
- Maksymilian Bilski (@maksbilski) - Backend Developer
- Sofiya Nasiakalia (@nasekajlo) - Frontend Developer

## License

MIT License
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
