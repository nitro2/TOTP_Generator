all: out

out: src/
	GO111MODULE=off GOPATH=${PWD} go build -o genotp ./src

clean:
	rm -f genotp