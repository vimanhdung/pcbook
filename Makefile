gen:
	protoc -I./proto proto/*.proto --go_out=plugins=grpc:pb

clean:
	rm -rf pb/*.go

run:
	go run main.go