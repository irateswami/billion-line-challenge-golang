run: 
	go run ./cmd/main.go

profiles:
	go tool pprof -pdf cpu.prof > cpu.pdf
#	go tool pprof -pdf block.prof > block.prof
#	go tool pprof -pdf threadcreate.prof > threadcreate.prof
#   go tool pprof -pdf goroutine.prof > goroutine.prof
	go tool pprof -pdf heap.prof > heap.prof

rmp: 
	rm *.pdf

