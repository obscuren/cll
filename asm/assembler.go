package asm

import (
	"container/list"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

func Assemble(ir *list.List) (asm []byte, err error) {
	for e := ir.Front(); e != nil; e = e.Next() {
		code := strings.Split(e.Value.(string), " ")
		switch len(code) {
		case 2:
			asm = append(asm, byte(vm.StringToOp(code[0])))
			asm = append(asm, common.String2Big(code[1]).Bytes()...)
		case 1:
			asm = append(asm, byte(vm.StringToOp(code[0])))
		default:
			return nil, fmt.Errorf("invalid IR %v", code)
		}
	}

	return
}
