package extensions

import (
	"github.com/silsuer/bingo_tpl"
	"strconv"
)

// 放入设定好的扩展，主动完成set和get的工作
type Ext_Core struct {
	bingo_tpl.Extension
}

// 添加操作符
func (ex *Ext_Core) GetName() string {
	return "core"
}

func (ex *Ext_Core) GetOperators() map[string]bingo_tpl.OperatorInterface {
	if !ex.IsInitialized() { // 如果没初始化过，先初始化
		ex.Init()
	}
	m := make(map[string]bingo_tpl.OperatorInterface)
	m["+"] = &bingo_tpl.Operator{Name: "+", Precedence: 30, Associativity: bingo_tpl.AssociativityLeft, Type: bingo_tpl.OperatorBinaryType, Target: func(tokens []bingo_tpl.Token) *bingo_tpl.Token {
		if len(tokens) != 2 {
			panic("the + operator is a binary operator.")
		}

		if tokens[0].T == bingo_tpl.TypeDigital && tokens[1].T == bingo_tpl.TypeDigital {
			r1, _ := strconv.Atoi(string(tokens[0].Value))
			r2, _ := strconv.Atoi(string(tokens[1].Value))
			r := r1 + r2
			t := new(bingo_tpl.Token)
			t.T = bingo_tpl.TypeDigital
			t.Value = []byte(strconv.Itoa(r))
			return t
		} else {
			t := new(bingo_tpl.Token)
			t.T = bingo_tpl.TypeString
			t.Value = []byte(string(tokens[0].Value) + string(tokens[1].Value))
			return t
		}
	}}

	m["-"] = &bingo_tpl.Operator{Name: "-", Precedence: 30, Associativity: bingo_tpl.AssociativityLeft, Type: bingo_tpl.OperatorBinaryType, Target: func(tokens []bingo_tpl.Token) *bingo_tpl.Token {
		if len(tokens) != 2 {
			panic("the - operator is a binary operator.")
		}

		if tokens[0].T == bingo_tpl.TypeDigital && tokens[1].T == bingo_tpl.TypeDigital {
			r1, _ := strconv.Atoi(string(tokens[0].Value))
			r2, _ := strconv.Atoi(string(tokens[1].Value))
			r := r1 - r2
			t := new(bingo_tpl.Token)
			t.T = bingo_tpl.TypeDigital
			t.Value = []byte(strconv.Itoa(r))
			return t
		} else {
			panic("the - operator needs 2 tokens of TypeDigital.")
		}
	}}

	return m
}
