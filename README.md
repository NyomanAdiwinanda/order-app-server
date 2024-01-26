# order-app-server

## Project setup

Before running this project, make sure [Golang](https://go.dev/doc/install) and [PostgreSQL](https://www.postgresql.org/download/) is already installed on your computer.

If already installed, follow all next steps below.

#### 1) Clone the repo

```
git clone https://github.com/NyomanAdiwinanda/order-app-server.git
```

#### 2) Configure Env File

inside the project repo, **rename** the file `.env.example` to `.env`

and inside that `.env` file, you will find something like this

```
DB_HOST=localhost
DB_PORT=5432
DB_NAME=order_app
DB_USER=
DB_PASSWORD=
```

Inside that `.env` file, assign the `DB_USER` with your **local postgres username** and the `DB_PASSWORD` with your **local postgres password**

#### 3) Running The Project

inside your terminal at this project directory, run:

```
go mod download
```

```
go run cmd/main.go
```

The project will run on localhost port 8080

#### 4) Running DB Migration

The database will need to be migrated first. To run database migration, go to `localhost:8080/migrate` inside your browser

It will return you a json like below if migration is success

![Screenshot 2024-01-27 at 06 08 34](https://github.com/NyomanAdiwinanda/order-app-server/assets/65802394/d8040d27-1329-40ed-9656-4e2616717398)

#### 5) Running DB Seeder

After migration, the database need to be seeded with the given CSV files from `/csv` folder.

To run DB seeder, go to `localhost:8080/seed` inside your browser

It will return you a json like below if seeder is success

![Screenshot 2024-01-27 at 06 15 59](https://github.com/NyomanAdiwinanda/order-app-server/assets/65802394/e049c049-816d-4063-9ff3-8978c3b7d3fe)

The project setup is now fully completed

## About the Project

The structure of this project is as given below

```bash
order-app-server/
│
├── /cmd
│   └── main.go          # The entry point of the program.
│
├── /internal
│   ├── /core
│   │   ├── /models      # Core business logic and data structures.
│   │   └── /usecases    # Application-specific business rules.
│   │
│   ├── /adapters
│   │   ├── /repositories   # Data access logic (interface implementations).
│   │   └── /handlers       # HTTP handlers (converts HTTP requests to domain logic calls).
│   │
│   └── /infrastructure
│       ├── /database       # Database initialization and connection logic.
│       └── /utils          # Utility functions and common helpers.
│
└── /csv    # For storing all the required CSV files.
```
