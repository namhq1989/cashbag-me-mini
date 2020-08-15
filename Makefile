export ZOOKEEPER_URI= 127.0.0.1:2181

run:
	go run main.go
test:
	go test ./... -v