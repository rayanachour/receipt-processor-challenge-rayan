                                                Receipt Processor API
A lightweight API service built in Go for processing receipts. This application accepts receipt data, calculates reward points, and returns a unique identifier for each processed receipt. You can retrieve the points associated with each receipt using its ID.

Table of Contents
    Running the Project
    API Endpoints
    Process Receipt
    Get Points
    Testing with Postman
    Example Usage
    Notes
Running the Project
    Install Go: Ensure Go is installed. If not, install Go.
    Start the Server:
    go run main.go
    This will start the server at http://localhost:8080, with output confirming the server is running.
    API Endpoints
1. Process Receipt
    URL: POST /receipts/process
    Description: Accepts a receipt in JSON format, calculates the reward points based on items and other criteria, and returns a unique ID for the receipt.
    Request Example:
    json
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
    Response Example:
    json
    {
      "id": "unique-receipt-id"
    }
2. Get Points
    URL: GET /receipts/{id}/points
    Description: Retrieves the calculated points for a specific receipt based on its ID.
    Response Example:
    json
    {
      "points": 20
    }
Testing with Postman
    Process a Receipt:

    Method: POST
        *URL: http://localhost:8080/receipts/process
        *Body: JSON (use the example provided above).
        *Get Points:

    Method: GET
        *URL: http://localhost:8080/receipts/{id}/points
        *Replace {id} with the receipt ID received from the Process Receipt endpoint.
Example Usage
    Submit a Receipt:
        *Send a POST request to /receipts/process with the receipt JSON data to get a unique receipt ID.
    Retrieve Points:
        *Use the ID from the previous response to send a GET request to /receipts/{id}/points to retrieve the associated points.
Notes
Data Persistence: This application uses in-memory storage, so all data is lost when the server is restarted. This setup is ideal for testing and development purposes but may need persistent storage for production environments.
