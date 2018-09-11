package bingo_tpl

// 方法
type TplFunc struct {
	name      string            // 方法名
	callback  interface{}       // 方法对应的回调
	options   map[string]string // 配置项
	arguments map[string]string // 参数
}

func (tf *TplFunc) GetName() string {
	return tf.name
}
