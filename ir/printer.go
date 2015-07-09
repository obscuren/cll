package ir

import (
	"container/list"
	"fmt"
)

func Print(asm *list.List) {
	i := 0
	for e := asm.Front(); e != nil; e = e.Next() {
		fmt.Printf("%6d   %s\n", i, e.Value)

		i++
	}
}

/* CLL IR

PUSH 10    ; PUSH1 10
PUSH 20    ; PUSH1 20
ADD        ; +
-> 0       ; PUSH4 DEST
JUMPI      ; JUMPI
...        ; code
@@ 0 @@    ; Jump dest

*/
