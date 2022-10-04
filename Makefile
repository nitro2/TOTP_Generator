all: out

out: src/
	GOPATH=${PWD} go build -o genotp ./src

clean:
	rm -f genotp