## Frontend Structure

```
├── src/
│   ├── app/                    # Next.js app directory
│   │   ├── api/               # API routes
│   │   ├── components/        # Reusable components
│   │   ├── contexts/          # React contexts
│   │   ├── (protected)/      # Protected routes
│   │   └── (public)/         # Public routes
│   ├── config/                # Configuration files
│   └── types/                 # TypeScript type definitions
```

## Environment Variables

Create a `.env` file in the root directory with the following variables:

```env
NEXTAUTH_SECRET=<your_secret_key>
NEXTAUTH_URL=http://<frontend_host>:<frontend_port>/
NEXT_PUBLIC_BACKEND_HOST=<backend_host>
NEXT_PUBLIC_BACKEND_PORT=<backend_port>
```

## Production Deployment

The project includes a `production.Dockerfile` for containerized deployment. To build and run the production container:

```bash
docker build -f production.Dockerfile -t band-manager .
docker run -p 3000:3000 band-manager
```

## Project Components

### Routing & Authentication

The application uses Next.js App Router with a clear separation between public and protected routes:

```
src/app/
├── (public)/            # Public routes - accessible without authentication
│   ├── login/          # Login page
│   ├── register/       # Registration page
│   └── page.tsx        # Landing page
│
└── (protected)/        # Protected routes - require authentication
    ├── dashboard/      # User dashboard
    ├── events/         # Event management
    ├── announcements/  # Announcements
    ├── manage/         # Management tools
    ├── subgroups/      # Subgroup management
    └── tracks/         # Track management
```

### Authentication

- Built with NextAuth.js
- Protected routes are wrapped in route groups using the (protected) directory
- Middleware (src/middleware.ts) automatically checks authentication for all routes in (protected)
- Middleware also verifies JWT tokens for all API requests
- Users attempting to access protected routes while unauthorized are redirected to the login page

### Features

- **Events**: Create and manage band events
- **Announcements**: Post and manage announcements
- **Tracks**: Manage music tracks and sheets
- **Subgroups**: Organize band members into subgroups
- **Manage**: Manage your band

### Components

- `Header`: Simple header component
- `Sidebar`: Sidebar with list of your bands
- `RequireManager/RequireModerator`: Role-based access control components
- `LoadingScreen`: Loading state component
- `Navigation`: Group navigation component with routing
