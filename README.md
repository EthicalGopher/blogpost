# BlogPost - Go-based RESTful Blog API

BlogPost is a robust, high-performance RESTful API for a blogging platform, built with **Go**, **Fiber**, and **GORM**. It features user authentication (JWT), post management, comments, categories, and tags.

## 🚀 Technology Stack

- **Language:** Go 1.25+
- **Framework:** [Fiber v3](https://gofiber.io/)
- **ORM:** [GORM](https://gorm.io/)
- **Database:** SQLite (default)
- **Authentication:** JWT (JSON Web Tokens)
- **Middleware:** Rate limiting, Custom JWT verification

## 🏗️ Project Architecture

The project follows a modular structure for better maintainability and separation of concerns:

- **`main.go`**: Entry point of the application.
- **`server/`**: Contains route handlers, server initialization, and middleware.
- **`models/`**: Defines the GORM database models.
- **`utils/`**: Helper functions for database operations (CRUD) and authentication.
- **`db/`**: Database connection and auto-migration logic.
- **`tests/`**: Independent test suite for API verification.
- **`repository/`**: Storage for the SQLite database file (`blog.db`).

## 🛠️ Getting Started

### Prerequisites

- Go installed (v1.25 or higher recommended)
- Git

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/blogpost.git
   cd blogpost
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run main.go
   ```
   The server will start listening on `http://localhost:8143`.

## 📂 Project Structure

```text
.
├── main.go            # Entry point
├── server/            # Fiber server and handlers
│   ├── main.go        # Router and server setup
│   ├── middleware.go  # JWT and Auth middleware
│   ├── post.go        # Post-related handlers
│   ├── user.go        # User-related handlers
│   ├── category.go    # Category handlers
│   └── tag.go         # Tag handlers
├── models/            # GORM models (User, Post, Comment, etc.)
├── utils/             # Database CRUD and auth helpers
├── db/                # Database connection and migration
├── tests/             # API tests
├── repository/        # SQLite database storage
├── go.mod             # Go modules file
└── go.sum             # Go checksums
```

## ✨ Key Features

- **User Management**: SignUp, SignIn, and Profile access.
- **Authentication**: Secure routes via JWT, including Refresh Token support.
- **Blog Posts**: Full CRUD for posts (Public view, Protected creation/update/delete).
- **Comments**: View and add comments to posts.
- **Categories & Tags**: Organize posts with a flexible category and tagging system.
- **Security**: Built-in rate limiting and route protection.

## 🧪 Testing

The project uses a Test-Driven Development (TDD) approach. Tests are located in the `tests/` directory and use an in-memory SQLite database for speed and isolation.

To run the tests:
```bash
go test ./tests/...
```

## 📝 Coding Standards

- **Go Idioms**: Strictly follows Go coding conventions.
- **Documentation**: All public and internal functions are documented with concise comments.
- **Surgical Updates**: Code changes are focused and maintain architectural integrity.

## 🤝 Contributing

1. Fork the project.
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`).
4. Push to the Branch (`git push origin feature/AmazingFeature`).
5. Open a Pull Request.

## 📄 License

MIT License - see the project for details.
