# DynamoDb Local Test

This project tests setting up a dynamo db locally and interacting with it



## Pre-requistes
- Docker
- Go
## Setting up instance for testing
1) To run the redis cluster follow steps mentioned here:
```
docker-compose up -d dynamodb
docker logs -f my-dynamodb
```

2) Access the Admin UI
```
DYNAMO_ENDPOINT=http://localhost:8000
dynamodb-admin
```
