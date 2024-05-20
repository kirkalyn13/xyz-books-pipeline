
# XYZ Books Pipeline

XYZ Books Pipeline to check and update incoming ISBNs from newly added books from the CRUD UI, and record new data to a CSV file.

## Prerequisites
- Golang 1.22
- XYZ App Server and Rabbit MQ Instance

## Usage

The initial implementation was to utilize message queues, to simulate instantaneous processing of new books from the CRUD UI App.
Upon adding a new book data, the data is sent to the RabbitMQ instance and is consumed by the pipeline, where it is processed to check the ISBNs and update as needed, then append the data to the CSV file. 
See installation notes for troubleshooting instructions related to this implementation.

To address the requirement of calling an API from the CRUD server, also implemented a continuous checker (triggered every 60 seconds for demo purposes),
which could also act as a redundancy check for already exisiting book data, to fetch book data with missing ISBNs then process as needed.

Thus, the pipeline primarily operates using the RabbitMQ instance, but should also work with/without the message queue.

## Installation

1. Clone repository.
2. Navigate to the root folder of the project.
3. Run `go mod tidy`.
4. Run `go build cmd/main.go`. This should generate the executable.
5. Run `./main` on the folder where the executable is located (preferrably the root folder of the project).
6. Alternatively, if you have Golang installed on your machine, you can simply run `go run cmd/main.go` from the project root folder.
7. The output CSV file can be found on `/xyz-books-pipeline/output/book-data.csv` .

**Note:** If the Rabbit MQ instance is up and running and messages are queued, the pipeline may need to restart to start receiving the queued messages.

## Rabbit MQ instance
- **URL** - http://localhost:15672/
- **Username** - guest
- **Password** - guest

## Author

- [Engr. Kirk Alyn Santos](https://github.com/kirkalyn13)
