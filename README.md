# Go Customer and Order Management Service

## Overview

This project is a simple service built using Go to manage customers and their orders through REST APIs. The service integrates authentication via OpenID Connect, and includes SMS notifications using Africa's Talking SMS gateway. The project also has unit tests with code coverage and is set up for CI/CD.

## Features

1. **Customer Management**: Add and manage customers with basic details like name, code, and phone number.
2. **Order Management**: Add orders with details such as item, amount, and timestamp.
3. **Authentication & Authorization**: Integrated OpenID Connect (OIDC) for secure access.
4. **SMS Notification**: When an order is created, the customer will receive an SMS alert.
5. **Unit Testing & CI/CD**: Includes unit tests with coverage reports and CI/CD configuration using GitHub Actions.

---

## Prerequisites

Ensure that you have the following before setting up the project:

- **Go version**: `>=1.23.1`
- **PostgreSQL**: A running instance for the customer and orders database.
- **Africa's Talking API**: Credentials for sending SMS. Use the sandbox environment for testing purposes.
- **OpenID Connect provider**: An OIDC provider (e.g., Google) for user authentication.
- **GitHub account**: For setting up CI/CD via GitHub Actions.

---

## Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/Stensis/savannah-go
cd savannah-go
```

### 2.Install Go Modules

```
go mod download
```

### 3. Setup Environment Variables

Create a .env file in the project root with the following keys:

```
DATABASE_URL=your_postgres_url
AFRICAS_TALKING_API_KEY=your_africas_talking_api_key
AFRICAS_TALKING_USERNAME=sandbox
AFRICAS_TALKING_ENVIRONMENT=sandbox
OIDC_CLIENT_ID=your_oidc_client_id
OIDC_CLIENT_SECRET=your_oidc_client_secret
OIDC_PROVIDER_URL=your_oidc_provider_url
```

### 4. Setup the Database

You can use the following SQL to create the customers and orders tables:

```
CREATE TABLE customers (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100),
  code VARCHAR(100),
  phone_number VARCHAR(15)
);

CREATE TABLE orders (
  id SERIAL PRIMARY KEY,
  customer_id INT REFERENCES customers(id),
  item VARCHAR(255),
  amount DECIMAL(10, 2),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

The service will be available on `http://localhost:8080.`

### My savannah psql database looks like this

  ![Alt text](/assets/psqldatabase.png "Creating a customer before auth")


# REST API Endpoints

### 1. Create Customer

- URL: `/customers`
- Method: `POST`
- Payload

```
{
    "name": "John Doe",
    "code": "C123",
    "phone_number": "+254712345678"
}
```

- Creating a customer Before auth.
  ![Alt text](/assets/invalidtoken.png "Creating a customer before auth")

- Creating a customer After auth.
  ![Alt text](/assets/usercreated.png "Creating a customer After auth")

### 2. Create Order

- URL: `/orders`
- Method: `POST`
- Payload

```
{
    "customer_id": 1,
    "item": "Laptop",
    "amount": 1500.00
}
```

- Creating Order related to a customer.
  ![Alt text](/assets/OrderSuccess.png "Creating a customer After auth")

### 3. OpenID Connect Callback

- URL: `/auth/callback`
- Method: `GET`
- Description: Handles OAuth2 token exchange for OpenID Connect.

* You get a prompt of your google email.
  ![Alt text](/assets/AuthLogin.png "Creating a customer After auth")

- Adding google Password.
  ![Alt text](/assets/Password.png "Creating a customer After auth")

- Lastly you get a 2-step verification prompt for safety.

- The token is created and auth is complete.
  ![Alt text](/assets/token.png "Creating a customer After auth")

# Africa's Talking SMS IntegrationAfrica's Talking SMS Integration

Once an order is created, an SMS alert is sent to the customer's phone number using Africa’s Talking API.

Make sure to set up your sandbox or live Africa's Talking account and configure the .env file accordingly.

## Africa's Talking Setup

1. Sign up for Africa's Talking.
2. Create a new app in the Africa's Talking sandbox.
3. Generate the API key and sandbox credentials.

- Sandbox example.
  ![Alt text](/assets/sms.png "Creating a customer After auth")
4. Use the sandbox API key and username `(sandbox)` for development and testing purposes.


- Prod sms example.
  ![Alt text](/assets/PRODsms.png "Creating a customer After auth")

# Unit Testing

This project uses Testify for writing unit tests and Go-SQLMock to mock the database interactions.

### Run Unit Tests:

```
go test ./... -v
```

### Code Coverage"

```
go test ./... -cover
```

## Continuous Integration (CI)

The project uses GitHub Actions for CI. The CI workflow is located at .github/workflows/ci.yml and is configured to:

- Install Go
- Run the unit tests with coverage
- Perform linting

## To set up CI:

1. Ensure your repository is hosted on GitHub.
2. GitHub Actions will automatically trigger based on the configuration provided in .github/workflows/ci.yml.

## Continuous Deployment (CD)

To deploy the service, you can configure any `PAAS/FAAS/IAAS` provider of your choice. For example, you can deploy to services like Heroku, AWS Lambda, or Google Cloud Run.

# Future Enhancements

- Add GraphQL Support: Replace the REST API with or augment it using GraphQL.
- Add Pagination: Implement pagination for customer and order listings.
- Add Logging: Implement more sophisticated logging with log levels and structured logging.
- CI/CD Improvements: Automate deployments to a cloud provider.# savannah-go
