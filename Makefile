build:
	cd application && rice embed-go && cd .. && \
	cd handlers && rice embed-go && cd .. && \
	go build
dev:
	rm -f `find . -name rice-box.go` && fresh
install:
	go get -u github.com/pilu/fresh
	go get -u github.com/GeertJohan/go.rice
	go get -u github.com/GeertJohan/go.rice/rice