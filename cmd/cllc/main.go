package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/codegangsta/cli"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/logger/glog"
	"github.com/obscuren/cll/asm"
	"github.com/obscuren/cll/ast"
	"github.com/obscuren/cll/ir"
)

var (
	app       *cli.App
	DebugFlag = cli.BoolFlag{
		Name:  "debug",
		Usage: "output full trace logs",
	}
)

func init() {
	app = utils.NewApp("0.0.1", "cll compiler")
	app.Flags = []cli.Flag{
		DebugFlag,
	}
	app.Action = run
}

func run(ctx *cli.Context) {
	debug := ctx.GlobalBool(DebugFlag.Name)

	if len(ctx.Args()) == 0 {
		glog.Exitln("err: <filename> required")
	}

	fn := ctx.Args().First()
	src, err := ioutil.ReadFile(fn)
	if err != nil {
		glog.Exitln("err:", err)
	}

	tree := ast.ParseFile(fn, string(src))

	if debug {
		fmt.Println("AST")
		ast.Print(tree)
	}

	intermediate, err := ir.ParseTree(tree)
	if err != nil {
		glog.Exitln(err)
	}
	if debug {
		fmt.Println("IR")
		ir.Print(intermediate)
	}

	bcode, err := asm.Assemble(intermediate)
	if err != nil {
		glog.Exitln(err)
	}

	fmt.Printf("%x\n", bcode)
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
