package bingo_tpl

const (
	OperatorUnaryType  = 1 // 一元运算符
	OperatorBinaryType = 2 // 二元运算符
	AssociativityLeft  = 1 // 左结合
	AssociativityRight = 2 // 右结合
)

// 操作符接口
type OperatorInterface interface {
	GetName() string    // 接口名
	GetPrecedence() int // 优先级
	GetTarget() interface{}
	GetOperatorType() int  // 1 是一元操作符 2是二元操作符
	GetAssociativity() int // 结合性
	//SetName(name string)     // 设置操作符名称
	//SetPrecedence(p int)     // 设置优先级
	//SetTarget(i interface{}) // 设置对应的方法
	//SetAssociativity(a int)  // 设置结合性
}

// 操作符
type Operator struct {
	Name          string               // 名称
	Precedence    int                  // 优先级
	Target        func(tokens []Token) *Token // 目标函数,传入一个token数组，返回一个token（左右子树作为参数，父节点作为）
	Associativity int                  // 结合性
	Type          int                  // 一元运算符还是二元运算符
}

func (op *Operator) GetName() string {
	return op.Name
}

func (op *Operator) GetPrecedence() int {
	return op.Precedence // 默认二元
}

func (op *Operator) GetTarget() interface{} {
	return op.Target
}

func (op *Operator) GetOperatorType() int {
	return op.Type
}

func (op *Operator) GetAssociativity() int {
	return op.Associativity
}
