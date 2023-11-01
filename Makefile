build:
	go build -o bin/tfmod ./src

compile:
	# Linux
	GOOS=linux GOARCH=arm go build -o bin/tfmod-linux-arm ./src
	GOOS=linux GOARCH=arm64 go build -o bin/tfmod-linux-arm64 ./src

	# Windows
	GOOS=windows GOARCH=amd64 go build -o bin/tfmod-windows-amd64.exe ./src
	GOOS=windows GOARCH=386 go build -o bin/tfmod-windows-386.exe ./src

	# Mac OS
	GOOS=darwin  GOARCH=amd64 go build -o bin/tfmod-amd64-darwin ./src