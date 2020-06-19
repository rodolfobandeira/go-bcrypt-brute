all: darwin linux windows

darwin:
	GOOS=darwin GOARCH=amd64 go build -o bin/go-bcrypt-brute go-bcrypt-brute.go

linux:
	GOOS=linux GOARCH=amd64 go build -o bin/go-bcrypt-brute-linux go-bcrypt-brute.go

windows:
	GOOS=windows GOARCH=amd64 go build -o bin/go-bcrypt-brute-windows.exe go-bcrypt-brute.go

clean:
	rm -f bin/go-bcrypt-brute bin/go-bcrypt-brute-linux bin/go-bcrypt-brute-windows.exe

