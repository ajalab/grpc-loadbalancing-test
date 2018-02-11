.PHONY: build containerize

SERVICE = lb-example-server:latest
TARGET = app
MAIN = main.go

build: $(MAIN)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(TARGET) $(MAIN)

containerize: Dockerfile $(TARGET)
	echo docker build -t $(SERVICE) .
	     docker build -t $(SERVICE) .
