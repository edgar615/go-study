package counter

//标识符以小写标识符开头时，这个标识符就是未公开的
//标识符以大写标识符开头时，这个标识符就是公开的
type alertCounter int

//New创建并返回一个未公开的alertCounter类型的值
func New(value int) alertCounter {
	return alertCounter(value)
}
