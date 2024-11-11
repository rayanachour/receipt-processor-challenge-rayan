## Running the project

1. [Install Go](https://go.dev/doc/install)
2. Run `go run main.go` in the root directory of the project
3. The server will start on `http://localhost:8080`. You should see output indicating the server has started.

## API Endpoints

### 1. Process Receipts

- **URL**: `POST /receipts/process`
- **Description**: Accepts a receipt JSON payload, calculates points, and returns a unique receipt ID.
- **Request Example**:

    ```json
    {
      "retailer": "Target",
      "purchaseDate": "2022-01-01",
      "purchaseTime": "13:01",
      "items": [
        {
          "shortDescription": "Mountain Dew 12PK",
          "price": "6.49"
        },
        {
          "shortDescription": "Emils Cheese Pizza",
          "price": "12.25"
        }
      ],
      "total": "35.35"
    }
    ```

- **Response Example**:

    ```json
    {
      "id": "unique-receipt-id"
    }
    ```

### 2. Get Points

- **URL**: `GET /receipts/{id}/points`
- **Description**: Retrieves the points for a given receipt ID.
- **Response Example**:

    ```json
    {
      "points": 20
    }
    ```

## Testing with Postman

1. **Process a Receipt**:
   - **Method**: `POST`
   - **URL**: `http://localhost:8080/receipts/process`
   - **Body**: JSON (example provided above).

2. **Get Points**:
   - **Method**: `GET`
   - **URL**: `http://localhost:8080/receipts/{id}/points`
   - Replace `{id}` with the ID received from the previous step.

## Example Usage

1. Submit a receipt via `POST /receipts/process` and note the ID returned.
2. Use the returned ID in a `GET /receipts/{id}/points` request to retrieve the points for that receipt.

## Notes

- **In-Memory Storage**: All data is stored in memory, meaning data will be lost if the server is restarted.