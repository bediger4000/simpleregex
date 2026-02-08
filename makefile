all: simpleregex matching

simpleregex: matching.go main.go
	go build $(PWD)

matching: cversion/matching.c
	cc -g -Wall -Wextra -o matching cversion/matching.c

clean:
	rm -rf simpleregex matching

test:
	go test -v  .
