main : cmd/main.go
	go build ./cmd/main.go
run : 
	./main
clean : 
	rm main
build :
	go build ./cmd/main.go
