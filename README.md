# Receipt Processor

## Overview

The Receipt Processor is a web service designed to process shopping receipts, calculate points based on a set of predefined rules, and retrieve the calculated points using an ID. The service is implemented in Go and uses in-memory storage, ensuring simplicity and ease of testing.

## Features
- Process Receipts
Endpoint to submit a receipt and calculate points based on predefined rules.

- Retrieve Points
Endpoint to retrieve the calculated points for a previously processed receipt.

- Stateless Data
Data is stored in memory and does not persist across application restarts.

## Installation and Running
### Prerequisites
- Install `Go`.
- Ensure `Docker` is installed (if running with `Docker`).

### Running Locally
1. Clone the repository:

```bash
git clone https://github.com/superdev808/receipt-processor-solution.git
cd receipt-processor-solution
```

2. Install dependencies and copy .env file:
```bash
go mod tidy
mv .env.example .env
```

3. Run the application:

```bash
go run .
```

### Running with Docker

1. Copy .env file

```bash
mv .env.example .env
```

2. Run Docker with Docker Compose:

```bash
docker compose up -d
```

3. Access the API at http://localhost:8080.

### `.env` File Configuration
Create a .env file in the project root to specify environment variables:
```env
PORT=8080
```

## Development Notes
### Code Organization:

- models/: Contains data structures like Receipt and Item.
- services/: Contains the business logic for processing receipts and calculating points.

### Testing:

- Add unit tests for CalculatePoints in services/receipt_service.go.
- Test endpoints using tools like Postman or curl.

### Data Persistence:

- This implementation uses in-memory storage for simplicity.
- For production, consider integrating a database like PostgreSQL or MongoDB.

### Running the Test
```bash
go test ./services -v
```

### Running Test Coverage
```bash
go test ./services -cover
```