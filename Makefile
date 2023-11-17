OUTPUT_DIR=./bin

NAME=queue-text-file

build: clean
	GOARCH=amd64 GOOS=darwin go build -trimpath -ldflags="-s -w" -o ${OUTPUT_DIR}/${NAME}-macos-amd64 ./cmd/${NAME}
	GOARCH=arm64 GOOS=darwin go build -trimpath -ldflags="-s -w" -o ${OUTPUT_DIR}/${NAME}-macos-arm64 ./cmd/${NAME}
	GOARCH=amd64 GOOS=linux go build -trimpath -ldflags="-s -w" -o ${OUTPUT_DIR}/${NAME}-linux-amd64 ./cmd/${NAME}
	GOARCH=arm64 GOOS=linux go build -trimpath -ldflags="-s -w" -o ${OUTPUT_DIR}/${NAME}-linux-arm64 ./cmd/${NAME}

clean:
	go clean
	rm -f ${OUTPUT_DIR}/${NAME}-*

test:
	go test -v ./...
