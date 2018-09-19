package bingo_tpl

// 扩展结合
type ExtensionSet struct {
	extensions         map[string]ExtensionInterface // 扩展
	initialized        bool                          // 是否已经初始化过
	runtimeInitialized bool                          // 是否在运行态中初始化过
	parsers            map[string]ParserInterface    // 解析器
	visitors           []string                      // 访问者
	filters            []string                      // 过滤器
	tests              []string                      // 测试
	functions          map[string]TplFunc            // 方法
	unaryOperators     []OperatorInterface           // 一元运算符
	binaryOperators    []OperatorInterface           // 二元运算符
	globals            map[string]interface{}        // 全局变量
	functionCallbacks  interface{}                   // 方法回调
	filterCallbacks    interface{}                   // 过滤器回调
	lastModified       int                           // 上次修改
}

func (eSet *ExtensionSet) Init() {
	eSet.initialized = true
	eSet.extensions = make(map[string]ExtensionInterface)
}

// 向扩展集中添加扩展
func (eSet *ExtensionSet) AddExtension(ext ExtensionInterface) {
	// 获取这个扩展的结构体名
	//t := reflect.TypeOf(ext).String()

	if eSet.initialized { // 已经初始化过，退出程序
		panic("Unable to register extension " + ext.GetName() + " as extensions have already been initialized.")
	}else{
		eSet.Init()
	}

	// 是否注册过
	if _, ok := eSet.extensions[ext.GetName()]; ok {
		panic("Unable to register extension " + ext.GetName() + " as it is already registered.")
	}

	// 赋值
	eSet.extensions[ext.GetName()] = ext
}

// 向集合中添加方法
func (eSet *ExtensionSet) AddFunction(f TplFunc) {

	if eSet.initialized { // 已经初始化过，退出程序
		panic("Unable to register extension " + f.GetName() + " as extensions have already been initialized.")
	}

	// 是否注册过
	if _, ok := eSet.extensions[f.GetName()]; ok {
		panic("Unable to register extension " + f.GetName() + " as it is already registered.")
	}

	// 加入扩展的方法中
	eSet.functions[f.GetName()] = f
}

// 得到这个扩展中注册的所有方法
func (eSet *ExtensionSet) GetFunctions() map[string]TplFunc {
	if !eSet.initialized {
		eSet.initExtensions()
	}
	return eSet.functions
}

// 初始化多个扩展
func (eSet *ExtensionSet) initExtensions() {
	// 初始化所有map
	for _, v := range eSet.extensions {
		eSet.initExtension(v)
	}
}

// 初始化扩展
func (eSet *ExtensionSet) initExtension(ext ExtensionInterface) {
	// 查看这个扩展中的function，filter，

	// 添加扩展中的方法
	for _, f := range ext.GetFunctions() {
		eSet.AddFunction(f)
	}
	// 添加扩展中的标记
	for _, t := range ext.GetTokenParsers() {
		eSet.parsers[t.GetTagName()] = t
	}

	// 添加节点访问器

	// 添加操作符
	for _, o := range ext.GetOperators() {
		// 一元操作符和二元操作符
		if o.GetOperatorType() == OperatorUnaryType {
			eSet.unaryOperators = append(eSet.unaryOperators, o)
		}

		if o.GetOperatorType() == OperatorBinaryType {
			eSet.binaryOperators = append(eSet.binaryOperators, o)
		}
	}

}

// 扩展接口
type ExtensionInterface interface {
	GetName() string // 获取扩展名
	Init()
	GetFunctions() map[string]TplFunc            // 获取所有注册过的方法
	GetTokenParsers() map[string]ParserInterface // 获取所有注册过的标记
	GetOperators() map[string]OperatorInterface  // 获取所有操作符
}

type Extension struct {
	name         string // 名称
	initialized  bool   // 是否执行过init方法
	functions    map[string]TplFunc
	tokenParsers map[string]ParserInterface
	operators    map[string]OperatorInterface
}

func (ex *Extension) Init() {
	if ex.initialized == false {
		ex.functions = make(map[string]TplFunc)
		ex.tokenParsers = make(map[string]ParserInterface)
		ex.operators = make(map[string]OperatorInterface)
	}
	ex.initialized = true
}

func (ex *Extension) GetName() string {
	if ex.initialized == false {
		ex.Init()
	}
	return ex.name
}

func (ex *Extension) GetFunctions() map[string]TplFunc {
	if ex.initialized == false {
		ex.Init()
	}
	return ex.functions
}

func (ex *Extension) GetTokenParsers() map[string]ParserInterface {
	if ex.initialized == false {
		ex.Init()
	}
	return ex.tokenParsers
}

func (ex *Extension) GetOperators() map[string]OperatorInterface {
	if ex.initialized == false {
		ex.Init()
	}

	return ex.operators
}
