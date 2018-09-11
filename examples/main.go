package main

import (
	"fmt"
)

// 打印一个模版的栗子

func main() {
	//loader := &bingo_tpl.Loader{Path: "tpl"}
	//env := bingo_tpl.NewEnv(loader)
	//env.Init()
	//fmt.Println(env)
	//env.OpenLexicalChain("tpl/hello.html")
	//env.OpenLexicalChain("/Users/silsuer/go/src/github.com/silsuer/bingo_tpl/examples/tpl/hello.html")
	//env.Tpl["/Users/silsuer/go/src/github.com/silsuer/bingo_tpl/examples/tpl/hello.html"].Print()
	//println(string(env.Tpl["tpl/hello.html"].Nodes[0].Content))
	//println(reflect.TypeOf(1).String())
	a := "hello)))world"
	for _,v:= range a{

		fmt.Println(v)
	}
}
