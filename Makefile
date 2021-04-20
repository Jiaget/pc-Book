gen:
	protoc --go-grpc_out=pb -I=proto proto/*.proto --go_out=pb 
clean:
	rm pb/*.go
test:
	go test -cover -race ./...
client:
	go run cmd/client/main.go -address 0.0.0.0:8080
server:
	go run cmd/server/main.go -port 8080
proxy:
	$Env:http_proxy="http://127.0.0.1:7890";$Env:https_proxy="http://127.0.0.1:7890"