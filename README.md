# 🐊 Gator

**Gator** is a lightweight command-line RSS feed aggregator written in Go. It allows users to register accounts, subscribe to RSS feeds, and aggregate blog posts into a PostgreSQL database.

---

## 📑 Table of Contents

- [Prerequisites](#-prerequisites)
- [Installation & Setup](#-installation--setup)
  - [1. Install & Configure PostgreSQL](#1-install--configure-postgresql)
  - [2. Install Goose & Run Migrations](#2-install-goose--run-migrations)
  - [3. (Optional) Install SQLC](#3-optional-install-sqlc)
  - [4. Configure Gator](#4-configure-gator)
  - [5. Build & Install Gator](#5-build--install-gator)
- [Usage & Commands](#-usage--commands)
  - [User Management](#user-management)
  - [Feed Management](#feed-management)
  - [Aggregation](#aggregation)
- [Configuration Reference](#-configuration-reference)

---

## 🛠 Prerequisites

Ensure you have the following installed on your system:

- [Go](https://go.dev/doc/install) (v1.23 or later recommended)
- [PostgreSQL](https://www.postgresql.org/) (v15 or later)
- [Goose](https://github.com/pressly/goose) (Database migration tool)
- [SQLC](https://sqlc.dev/) (Type-safe SQL compiler for Go)

---

## 🚀 Installation & Setup

### 1. Install & Configure PostgreSQL

#### **Installation**

- **macOS (via Homebrew):**

  ```bash
  brew install postgresql@15
  ```

- **Linux / WSL (Debian/Ubuntu):**

  ```bash
  sudo apt update
  sudo apt install postgresql postgresql-contrib
  ```

#### **Verify Installation & Start Service**

Verify that PostgreSQL 15+ is installed:

```bash
psql --version
```

Start the PostgreSQL service:

- **macOS:**

  ```bash
  brew services start postgresql@15
  ```

- **Linux / WSL:**

  ```bash
  sudo service postgresql start
  ```

*(Linux / WSL only)* Set the system `postgres` user password:

```bash
sudo passwd postgres
```

#### **Create Database & User**

1. Access the `psql` console:
   - **macOS:**

     ```bash
     psql postgres
     ```

   - **Linux / WSL:**

     ```bash
     sudo -u postgres psql
     ```

2. Create the `gator` database and configure the user password:

   ```sql
   CREATE DATABASE gator;
   \c gator
   ALTER USER postgres PASSWORD 'postgres';
   ```

3. Exit the `psql` console:

   ```text
   exit
   ```

---

### 2. Install Goose & Run Migrations

Install the latest version of [Goose](https://github.com/pressly/goose):

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Verify the installation:

```bash
goose -version
```

#### **Connection String Format**

```text
postgres://<username>:<password>@<host>:<port>/<database>
```

- **macOS (default):** `postgres://<username>:@localhost:5432/gator`
- **Linux (default):** `postgres://postgres:postgres@localhost:5432/gator`

#### **Apply Database Migrations**

Navigate to the `sql/schema` directory and run the `up` migration:

```bash
cd sql/schema
goose postgres "postgres://postgres:postgres@localhost:5432/gator" up
```

Verify tables were created:

```bash
psql gator -c "\dt"
```

> [!TIP]
> To rollback migrations if needed, run:
>
> ```bash
> goose postgres "postgres://postgres:postgres@localhost:5432/gator" down
> ```

---

### 3. (Optional) Install SQLC

SQLC compiles SQL queries into type-safe Go code.

Install SQLC:

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

Verify installation:

```bash
sqlc version
```

To re-generate Go code from `sql/queries` and `sql/schema`:

```bash
sqlc generate
```

---

### 4. Configure Gator

Create a configuration file in your home directory at `~/.gatorconfig.json`:

```json
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

> [!IMPORTANT]
> Make sure to append `?sslmode=disable` to your connection string in `~/.gatorconfig.json` for local development.

---

### 5. Build & Install Gator

Install the CLI globally to your `$GOPATH/bin`:

```bash
go install .
```

Or build a local executable binary:

```bash
go build -o gator .
```

---

## 💻 Usage & Commands

Run Gator using the `gator` command (or `./gator` if using a local binary):

### User Management

| Command | Arguments | Description |
| :--- | :--- | :--- |
| `register` | `<name>` | Register a new user and set as the current active user |
| `login` | `<name>` | Set the current active user |
| `users` | — | List all registered users |
| `reset` | — | Reset database tables (wipes users & feeds) |

**Examples:**

```bash
gator register alice
gator login alice
gator users
```

### Feed Management

| Command | Arguments | Description |
| :--- | :--- | :--- |
| `addfeed` | `<name> <url>` | Add a new RSS feed for the current user |
| `feeds` | — | List all feeds stored in the database |
| `follow` | `<url>` | Follow an existing feed |

**Examples:**

```bash
gator addfeed "Boot.dev Blog" "https://blog.boot.dev/index.xml"
gator feeds
gator follow "https://blog.boot.dev/index.xml"
```

### Aggregation

| Command | Arguments | Description |
| :--- | :--- | :--- |
| `agg` | `<time_between_reqs>` | Run the feed scraper continuously at a duration accepted by Go's `time.ParseDuration` (e.g. `1s`, `30s`, `1m`, `1h`) |

**Example:**

```bash
gator agg 1m
```

---

## ⚙️ Configuration Reference

The `~/.gatorconfig.json` file supports the following fields:

| Field | Type | Description |
| :--- | :--- | :--- |
| `db_url` | `string` | PostgreSQL connection URL with `?sslmode=disable` |
| `current_user_name` | `string` | Name of the currently logged-in user |
