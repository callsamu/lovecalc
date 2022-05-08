build: 
	go build -o bin/web cmd/web/*

run: build
	bin/./web
