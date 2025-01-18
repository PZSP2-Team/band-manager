### Architectural Layers

#### Handlers (`/internal/handlers`)
Handles HTTP requests and responses. This layer is responsible for:
- Request parsing and validation
- Response formatting
- Route handling
- HTTP-specific logic
- Converting HTTP requests into usecase calls

#### Usecases (`/internal/usecases`)
Contains the business logic of the application. This layer:
- Implements core business rules
- Orchestrates data flow between entities
- Handles authorization logic
- Coordinates between different repositories and services
- Is independent of HTTP or database concerns

#### Repositories (`/internal/repositories`)
Manages data persistence and database operations. This layer:
- Handles all database interactions
- Implements CRUD operations
- Uses GORM for database operations
- Abstracts database implementation details from the rest of the application

#### Models (`/internal/model`)
Contains database model structures that:
- Define the database schema using GORM tags
- Map to database tables
- Define relationships between different entities
- Handle database-specific annotations and validations

#### Services (`/internal/services`)
Handles external integrations and infrastructure concerns:
- Email service integration
- Google Calendar integration
- Database connection management
- External API interactions
- Infrastructure-related functionalities

#### Domain (`/internal/domain`)
Contains pure business domain structures that:
- Define core business types
- Are independent of database or external concerns
- Represent business concepts and rules
- Are used across different layers of the application
