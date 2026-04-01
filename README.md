# AP2AS Distributed Order & Payment System

A distributed system consisting of two Go microservices for managing orders and processing payments.

## Architecture

- **Order Service**: Handles order lifecycle (Creation, Retrieval, Cancellation).
- **Payment Service**: Handles payment authorization and transaction logging.
- **Database**: Each service has its own PostgreSQL database.
- **Communication**: Order service communicates with Payment service via synchronous HTTP calls with a 2-second timeout.

## Services

### Order Service (`order-service`)
- **Default Port**: `8080`
- **Endpoints**:
  - `POST /orders/`: Create a new order (Status: `Pending` -> `Paid`/`Failed`).
  - `GET /orders/:id`: Get order details.
  - `PATCH /orders/:id/cancel`: Cancel a `Pending` order.

### Payment Service (`payment-service`)
- **Default Port**: `8081`
- **Endpoints**:
  - `POST /payments/`: Authorize a payment (Declines if amount > 100,000).
  - `GET /payments/:order_id`: Get payment status for a specific order.

## Setup

1. **Configure Environment**:
   - Copy `.env.example` to `.env` in both service directories.
   - Update database credentials and service URLs.

2. **Database Migrations**:
   - Run migrations for both services using your preferred tool (e.g., `golang-migrate`).

3. **Run Services**:
   ```bash
   cd order-service && go run cmd/main.go
   cd payment-service && go run cmd/main.go
   ```

## Business Rules

- **Financial Accuracy**: All monetary amounts are handled as `int64` (cents).
- **Immutability**: Once an order is paid, it cannot be cancelled.
- **Payment Limits**: Single payments over 1,000 units (100,000 cents) are automatically declined.
- **Timeouts**: Order service enforces a strict 2-second timeout when calling the Payment service.
