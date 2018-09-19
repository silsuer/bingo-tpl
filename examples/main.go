package main

import (
	"github.com/silsuer/bingo_tpl"
	"fmt"
	"github.com/silsuer/bingo_tpl/extensions"
)

// 打印一个模版的栗子

func main() {
	loader := &bingo_tpl.Loader{Path: "/Users/silsuer/go/src/github.com/silsuer/bingo_tpl/examples/tpl"}
	env := bingo_tpl.NewEnv(loader, make(map[string]string))
	env.AddExtension(new(extensions.Ext_Core))
	//env.Init()
	//fmt.Println(env)

	// 加载模版的过程，就是将模版的数据，解析成token流
	a := env.LoadTemplate("hello")
	fmt.Println(a)
	//fmt.Println(t.GetContentSource())
	//env.OpenLexicalChain("/Users/silsuer/go/src/github.com/silsuer/bingo_tpl/examples/tpl/hello.html")
	//env.OpenLexicalChain("hello.html")
	//env.OpenLexicalChain("/Users/silsuer/go/src/github.com/silsuer/bingo_tpl/examples/tpl/hello.html")
	//env.Tpl["/Users/silsuer/go/src/github.com/silsuer/bingo_tpl/examples/tpl/hello.html"].Print()
	//println(string(env.Tpl["tpl/hello.html"].Nodes[0].Content))
	//println(reflect.TypeOf(1).String())
	//a := "hello)))world\n\n\n+-*/^&|\\"
	//for _, v := range a {
	//	fmt.Println(v)
	//}
}
