.PHONY: server

BINARY_NAME=ko
INSTALL_DIR=/usr/local/bin
CUSTOM_BINARY_FOLDER=ko

build:
	go build -o $(BINARY_NAME) ./cli

install: build
	cp $(BINARY_NAME) $(INSTALL_DIR)/$(CUSTOM_BINARY_FOLDER)/$(BINARY_NAME)

publish: clean build install

clean:
	rm -f $(BINARY_NAME)

server:
	go run ./server
