gen:
	protoc --go-grpc_out=pb -I=proto proto/*.proto --go_out=pb 
clean:
	rm pb/*.go
test:
	go test -cover -race ./...
proxy:
	$Env:http_proxy="http://127.0.0.1:7890";$Env:https_proxy="http://127.0.0.1:7890"