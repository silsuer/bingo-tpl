package bingo_tpl

const (
	OperatorUnaryType  = 1 // 一元运算符
	OperatorBinaryType = 2 // 二元运算符
)

// 操作符接口
type OperatorInterface interface {
	GetName() string    // 接口名
	GetPrecedence() int // 优先级
	GetTarget() interface{}
	GetOperatorType() int            // 1 是一元操作符 2是二元操作符
	GetAssociativity() int   // 结合性
	SetName(name string)     // 设置操作符名称
	SetPrecedence(p int)     // 设置优先级
	SetTarget(i interface{}) // 设置对应的方法
	SetAssociativity(a int)  // 设置结合性
}

// 操作符
type Operator struct {
	name          string      // 名称
	precedence    int         // 优先级
	target        interface{} // 目标函数
	associativity int         // 结合性
}
