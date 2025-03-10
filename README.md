# Go protobuf namespace conflict
This repository demonstrates a protobuf namespace conflict. 

## Run
- run `go run main.go` to see the conflict

## Cause of the conflict
The conflict results from two different `*.pb.go` files having the same package name. More on namespace conflicts can be found [here](https://protobuf.dev/reference/go/faq/#namespace-conflict).