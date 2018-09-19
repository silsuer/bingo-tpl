package bingo_tpl

type Template struct {
	source []byte        // 源代码
	chain  *LexicalChain // 这个模版的词法链
	stream *TokenStream  // token流
	root   *ParserNode   // 抽象语法树根节点
	size   int           // 文件大小
}

func (t *Template) GetContentSource() string {
	return string(t.source)
}

//func tokenize()   {
//
//}
