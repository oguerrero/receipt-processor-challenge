## Receipt processor challenge - Fetch

***

You have two ways to run this server:

1. Build the project and run the executable file

Run go build in the root of the project to generate the .exe file

```bash
go build
```

Double click on the generated .exe file to run the server, make sure to allow the program to access the network.

2. Run the project directly from the source code

Use go run in the root of the project to create the server.

```bash
go run .
```

***

### Endpoints

#### POST /receipts

`http://localhost:8080/receipts/process`

This endpoint receives a JSON file with the receipts to be processed.

#### GET /receipts

`http://localhost:8080/receipts/ac8b3275-43a2-4798-9ce4-f9a3733285b0/points`

This endpoint returns a JSON file with the receipts that have been processed.

### Postman - Thunder Client

You can import the `thunderclient.json` or `postman.json` file to test the endpoints in their respective applications.

### Using curl

Testing the POST endpoint

```bash
curl --location 'http://localhost:8080/receipts/process' \
--header 'Content-Type: application/json' \
--data '{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}'
```

Testing the GET endpoint

```bash
curl --location 'http://localhost:8080/receipts/<receipt-uuid>/points'
```

### Running controller tests

You can run the controller tests with the following command:

```bash
go test ./test
```