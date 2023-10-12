default:
	GOOS=darwin GOARCH=arm64 go build -o bin/convert_macos-arm
	GOOS=darwin GOARCH=amd64 go build -o bin/convert_macos-amd64
	GOOS=windows GOARCH=arm64 go build -o bin/convert_windows-arm.exe
	GOOS=windows GOARCH=amd64 go build -o bin/convert_windows-amd64.exe
	GOOS=linux GOARCH=arm64 go build -o bin/convert_linux-arm
	GOOS=linux GOARCH=amd64 go build -o bin/convert_linux-amd64
	GOOS=linux GOARCH=arm GOARM=5 go build -o bin/convert_linux-rpi
