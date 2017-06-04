package search

// defaultMatcher implements the default matcher.
//空结构在创建实例时，不会分配任何内存，这种结构很适合创建没有任何状态的类型
type defaultMatcher struct{}

// init registers the default matcher with the program.
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search implements the behavior for the default matcher.
//实现match.go中定义对Matcher接口的行为
//方法声明为使用defaultMatcher类型的值作为接收者
//如果声明函数的时候带有接收者，意味着声明了一个方法，这个方法会和指定的接收者的类型绑定在一起
//这意味着我们可以使用defaultMatcher类型的值或者指向这个类型值的指针来调用Search方法
//这意味着我们可以使用接收者类型的值来调用这个方法，还是使用接收者类型值的指针来调用这个方法,编译器都会正确地引用或者解引用对应对值，作为接收者传递给Search方法
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}

/*
调用方法的例子
1.方法声明为使用defaultMatcher类型的值作为接收者
func (m defaultMatcher) Search(feed *Feed, searchTerm string)
声明一个指向defaultMatcher类型值的指针
dm := new(defaultMatcher)
编译器会解开dm指针的引用，使用对应的值调用方法
dm.Search(feed, "test")

2.方法声明为使用指向defaultMatcher类型值的指针作为接收者
func (m *defaultMatcher) Search(feed *Feed, searchTerm string)
声明一个defaultMatcher的类型值
var dm defaultMatcher
编译器会自动生成指针引用dm值，使用指针调用方法
dm.Search(feed, "test")

因为大部分方法在调用之后都需要维护接收者的值的状态，所以一个最佳实践是，将方法的接收者声明为指针
对于defaultMatcher类型来说，使用值作为接收者是因为创建一个defaultMatcher类型的值不需要分配内存．
由于defaultMatcher不需要维护状态，所以不需要指针形式的接收者

与直接通过值或者指针调用方法不同，如果通过接口类型的值调用方法，规则有很大不同
使用指针作为接收者声明对方法，只能在接口类型的值是一个指针的时候被调用．
使用值作为接收者声明的方法，在接口类型的值为值或者指针时都可以调用

func (m *defaultMatcher) Search(feed *Feed, searchTerm string)
通过interface的值来调用方法
var dm defaultMatcher
var matcher Matcher = dm//将值赋给接口类型
matcher.Search(feed, "test")
//error

func (m defaultMatcher) Search(feed *Feed, searchTerm string)
var dm defaultMatcher
var matcher Matcher = &dm//将指针赋给接口类型
matcher.Search(feed, "test")
//successful
*/
