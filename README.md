# fundraising-backend-api

This is the backend for the Fundraising application, which allows users to create and manage fundraising campaigns. It is built with Golang, using the Gin framework, Gorm ORM, JWT for authentication, and Midtrans for payment processing.

## Features

- User Authentication (Registration, Login, JWT-based)
- Campaign Management (Create, Update, Delete, View)
- Campaign Image Management
- Transaction Processing and Tracking
- Integration with Midtrans for payment processing

## ERD (Entity-Relationship Diagram)

![image](https://github.com/user-attachments/assets/85599426-f99c-4bad-b5fc-175949139d33)

## Database Tables

### Users
- **id**: int (Primary Key)
- **name**: varchar
- **occupation**: varchar
- **email**: varchar
- **password_hash**: varchar
- **avatar_file_name**: varchar
- **role**: varchar
- **token**: varchar
- **created_at**: datetime
- **updated_at**: datetime

### Campaigns
- **id**: int (Primary Key)
- **user_id**: int (Foreign Key referencing `Users`)
- **name**: varchar
- **short_description**: varchar
- **description**: text
- **goal_amount**: int
- **current_amount**: int
- **perks**: text
- **backer_count**: int
- **slug**: varchar
- **created_at**: datetime
- **updated_at**: datetime

### Campaign Images
- **id**: int (Primary Key)
- **campaign_id**: int (Foreign Key referencing `Campaigns`)
- **file_name**: varchar
- **is_primary**: boolean (tinyint)
- **created_at**: datetime
- **updated_at**: datetime

### Transactions
- **id**: int (Primary Key)
- **campaign_id**: int (Foreign Key referencing `Campaigns`)
- **user_id**: int (Foreign Key referencing `Users`)
- **amount**: int
- **status**: varchar
- **code**: varchar
- **created_at**: datetime
- **updated_at**: datetime

## Getting Started

### Prerequisites
- Golang 1.18+
- MySQL or PostgreSQL
- Midtrans account for payment integration

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/fundraising-api.git
   cd fundraising-api
   ```

2. Install dependencies:
   ```
    go mod tidy
   ```

3. Set up your environment variables. Create a .env file in the root directory and configure the following:
   ```
    DB_DRIVER=pgsql
    DB_HOST=127.0.0.1
    DB_PORT=5432
    DB_USER=root
    DB_PASSWORD=yourpassword
    DB_NAME=fundraising_db
    
    JWT_SECRET=your_jwt_secret
    
    MIDTRANS_SERVER_KEY=your_midtrans_server_key
    MIDTRANS_CLIENT_KEY=your_midtrans_client_key
   ```
4. Run database migrations:
  ```
    go run main.go migrate
  ```
5. Start the server:
  ```
    go run main.go
  ```

## Getting Started
- POST /register: Register a new user.
- POST /login: Authenticate a user and receive a JWT.
- GET /campaigns: List all campaigns.
- POST /campaigns: Create a new campaign (requires authentication).
- GET /campaigns/:id: Get details of a campaign.
- PUT /campaigns/:id: Update a campaign (requires authentication).
- DELETE /campaigns/:id: Delete a campaign (requires authentication).
- POST /transactions: Create a transaction (requires authentication).
- GET /transactions/:id: Get details of a transaction.
