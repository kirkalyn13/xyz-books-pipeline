
# XYZ Books Pipeline

XYZ Books Pipeline to check and update incoming ISBNs from newly added books from the CRUD UI, and record new data to a CSV file.

## Prerequisites
- Golang

## Installation

1. Clone repository.
2. Navigate to the root folder of the project.
3. Run `go build cmd/main.go`. This should generate the executable.

## Deployment

1. Run `./main` on the folder where the executable is located (preferrably the root folder of the project).
2. The output CSV file can be found on `/xyz-books-pipeline/output/book-data.csv` .

**Note:** If the Rabbit MQ instance is up and running and messages are queued, the pipeline may need to restart to start receiving the queued messages.

## Rabbit MQ instance
- **URL** - http://localhost:15672/
- **Username** - guest
- **Password** - guest

## Author

- [Engr. Kirk Alyn Santos](https://github.com/kirkalyn13)
