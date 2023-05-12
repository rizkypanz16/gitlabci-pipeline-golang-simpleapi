## How to run golang app

1. Initialize a new module
`go mod init goproject`

2. Install the GO dependencies
`go get -u fmt encoding/json log net/http`

3. Build the Go app
`go build -o binary .`

4. Run the executable
`./binary`