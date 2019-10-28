run:
	echo "Running.."
	go run -race ./cmd/counter-api/main.go

build:
	echo "Boulding.."
	go build -o bin/counter-api ./cmd/counter-api/main.go

test:
	echo "Testing.."
	go test -count=1 -p=1 -v ./...

compile:
	echo "Compiling for Mac OS.."
	GOOS=darwin GOARCH=amd64 go build -o bin/counter-api-darwin-amd64 ./cmd/counter-api/main.go
	