test:
	cd src; go test ./... -v

build:
	mkdir -p out; cd src;  go build -o ../out/simpleapi cmd/main.go

clean:
	rm -rf out

.PHONY: clean build test