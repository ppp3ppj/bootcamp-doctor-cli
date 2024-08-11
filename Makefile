build:
	go build -o .build/bootcamp-doctor-cli main.go

run:
	go run main.go

compile:
	echo "Compling for every OS/platform"
	GOOS=linux GOARCH=arm go build -o build/bootcamp-doctor-cli-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o build/bootcamp-doctor-cli-linux-arm64 main.go
	GOOS=linux GOARCH=amd64 go build -o build/bootcamp-doctor-cli-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o build/bootcamp-doctor-cli-windows-amd64.exe main.go
