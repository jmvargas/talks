test:
	go test -v ./...

benchmark:
	go test -cpuprofile cpu.prof -bench .
	
results:
	pprof -http=localhost:8080 cpu.prof