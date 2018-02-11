.PHONY: build containerize

SERVICE = lb-example-client
TARGET = app
MAIN = main.go
TAG = 3

build: $(MAIN)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(TARGET) $(MAIN)

containerize: Dockerfile $(TARGET)
	echo docker build -t $(SERVICE):$(TAG) .
	     docker build -t $(SERVICE):$(TAG) .
