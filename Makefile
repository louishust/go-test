build:
	go build -o cobra cobra.go
	go build -o pipeline pipeline.go
	go build -o context context.go

clean:
	-rm -f cobra pipeline context
