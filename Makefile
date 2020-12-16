include config.mk

findmt: findmt.go
	@go build

install:
	@mkdir -p ${PREFIX}/bin
	@install findmt ${PREFIX}/bin

uninstall:
	@rm -f ${PREFIX}/bin/findmt

clean:
	@rm -f findmt

test:
	go test

