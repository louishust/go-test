build:
	go build -o cobra cobra.go
	go build -o pipeline pipeline.go
	go build -o context context.go
	go build -o context-timeout context-timeout.go
	go build -o channel-test channel-test.go

clean:
	-rm -f cobra pipeline context context-timeout
	-rm -f channel-test
