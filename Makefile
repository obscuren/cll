all:
	go tool yacc -o ast/gll.go ast/gll.y
	go install ./cmd/gllc
	gllc --debug ./examples/basic.cll

