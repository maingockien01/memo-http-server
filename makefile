build:
	go mod tidy
	go build ./main.go

runServer:
	go mod tidy
	go run -a -v ./main.go

buildScaper:
	cc -Wall -Wextra -Wpedantic -g ./scaper.c -o scaper

runScaper:
	./scaper

clean:
	rm -f scaper main

