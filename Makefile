build : cmd/main.go
	go build -o main ./cmd/main.go
run : 
	./main
clean : 
	rm main