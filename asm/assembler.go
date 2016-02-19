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

			if len(code[1]) > 1 && code[1][:2] == "0x" {
				asm = append(asm, common.FromHex(code[1])...)
			} else {
				num := common.String2Big(code[1]).Bytes()
				if len(num) == 0 {
					num = []byte{0}
				}
				asm = append(asm, num...)
			}
		case 1:
			asm = append(asm, byte(vm.StringToOp(code[0])))
		default:
			return nil, fmt.Errorf("invalid IR %v", code)
		}
	}

	return
}
