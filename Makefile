build:
	go build -o cobra cobra.go
	go build -o pipeline pipeline.go
	go build -o context context.go
	go build -o context-timeout context-timeout.go

clean:
	-rm -f cobra pipeline context context-timeout
