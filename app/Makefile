all:
	clang++ main.cc -o main -L. -lhello
lib:
	clang++ -Wall -fPIC -shared -g hello.c -o libhello.so

app:
	go build -o wedding_app main1.go sql1.go test.go
