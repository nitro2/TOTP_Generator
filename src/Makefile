all: out

out: main.go otp/
	GOPATH=${PWD}/.. go build -o genotp

clean:
	rm -f genotp