# Go Delivery Service - Kiwibot

This is an API service built with Golang+Gin+Firebase, to solve the Kiwibot [Backend Challenge](https://kiwi.notion.site/Backed-developer-technical-test-bbea8f94184643419b57932d214ed66f)

## Pre requirements

- Go 1.18+
- Firebase database (cloud firestore)

## Instructions to run in a Local environment

1. Create an `.envrc` file with the variables of the `.envrc.example` and set your environment values.
    - When create the firebase app, download the service-account.json file and the path of the file is the one that must set as value of `GOOGLE_FIREBASE_CREDENTIALS_PATH`
    ```
    export PORT=8080 # an available port in your machine.
    export GIN_MODE=debug # debug, release (last one for production)
    export GOOGLE_FIREBASE_CREDENTIALS_PATH="/path/to/service-account.json"
    ```

2. Load the .envrc file. In Mac/Linux run: `source .envrc`

3. Run the service. `go run main.go`

### Aditional notes

The swagger docs could be find in http://localhost:8080/docs/index.html (if needed, change the port as the same that you have in the .envrc file)

To see the coverage of the tests made, you can run go tool `go tool cover -html=cover.out`

In the [docs](./docs) folder you could found a postman collection that you can import in Postman Desktop or Web to test the endpoints created.

### Author

- Maria Mora <mariamora2807@gmail.com>