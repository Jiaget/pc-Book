gen:
	protoc -I=proto proto/*.proto --go_out=pb
clean:
	rm pb/*.go
test:
	go test -cover -race ./...