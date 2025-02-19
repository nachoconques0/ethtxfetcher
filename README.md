# ethtxfetcher

### What I've done
A quick API that has a controller & service layer. To the service we are injecting a client that will call via HTTP the JSON-RPC ETH API.

- Client interacts with JSONRPC Ethereum Blockchain
- We can Subscribre an address
- We can Get latest block
- We can Get Txs of the latest block if the address had interacted with it

### How to run it :scream_cat:
1. `git clone` this repo
2. Once you are inside the repo make sure to have the modules if some issues appears just run "go mod tidy" or "make mod"
3. Then run make mod


### Exposed HTTP Endpoints
- `GET /block`
- `POST /subscribe/{address}`
- `GET /address/{address}/tx`   

### Extra

Sorry for the quality of the project. Im in sort of vacation and also I haven't touched any coding since I ended my last job. Which was around 4 months!

### File Directory

```
📦ethtxfetcher
 ┣ 📂cmd
 ┃ ┗ 📂server
 ┃ ┃ ┣ 📜dev.go
 ┃ ┃ ┗ 📜main.go
 ┣ 📂internal
 ┃ ┣ 📂api
 ┃ ┃ ┗ 📜server.go
 ┃ ┣ 📂app
 ┃ ┃ ┣ 📜app.go
 ┃ ┃ ┣ 📜domain.go
 ┃ ┃ ┗ 📜options.go
 ┃ ┣ 📂controller
 ┃ ┃ ┗ 📜parser.go
 ┃ ┣ 📂errors
 ┃ ┃ ┗ 📜errors.go
 ┃ ┣ 📂model
 ┃ ┃ ┗ 📜parser.go
 ┃ ┣ 📂parser
 ┃ ┃ ┗ 📜parser.go
 ┃ ┗ 📂service
 ┃ ┃ ┗ 📜parser.go
 ┣ 📜.gitignore
 ┣ 📜Makefile
 ┣ 📜README.md
 ┣ 📜go.mod
 ┣ 📜go.sum
 ┗ 📜postman.json
 ```