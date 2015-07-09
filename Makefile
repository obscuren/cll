all:
	go tool yacc -o ast/cll.go ast/cll.y
	go install ./cmd/cllc
	cllc --debug ./examples/basic.cll

