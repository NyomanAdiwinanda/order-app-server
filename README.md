```bash
your_project/
│
├── /cmd
│   └── main.go          # The entry point of the program.
│
├── /internal
│   ├── /core
│   │   ├── /entities    # Core business logic and data structures.
│   │   └── /usecases    # Application-specific business rules.
│   │
│   ├── /adapters
│   │   ├── /repositories    # Data access logic (interface implementations).
│   │   └── /handlers       # HTTP handlers (converts HTTP requests to domain logic calls).
│   │
│   └── /infrastructure
│       ├── /database       # Database initialization and connection logic.
│       └── /utils          # Utility functions and common helpers.
│
└── /pkg
    └── /api               # Interface definitions for repositories, handlers, etc.
```
