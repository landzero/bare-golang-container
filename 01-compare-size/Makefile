all: hello-c hello-c-static hello-go

hello-c: hello.c
	gcc -g -o hello-c hello.c

hello-c-static: hello.c
	gcc -g -static -static-libgcc -o hello-c-static hello.c

hello-go: hello.go
	go build -o hello-go hello.go
