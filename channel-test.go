package main

func main() {
	serveSQL := make(chan bool)
	close(serveSQL)

	select {
	case <-serveSQL:
	}

}
