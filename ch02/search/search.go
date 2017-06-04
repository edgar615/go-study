package search

import (
	"log"
	"sync"
)

// A map of registered matchers for searching.
//包级变量，声明为Matcher类型的映射(map)，这个映射以string类型值作为键，Matcher类型值作为映射后对值
//make函数初始化类map
//在Go中，所有的语言都被初始化为其零值，数值类型：0，字符串：空字符串，bool：false，指针：nil
//对于引用类型来说，所引用的底层数据结构会初始化为对应对零值．但是被声明为零值的引用类型的变量，会返回nil作为其值
var matchers = make(map[string]Matcher)

// Run performs the search logic.
//func用来声明函数，func关键字后面紧跟着函数名，参数，以及返回值
func Run(searchTerm string) {
	// Retrieve the list of feeds to search through.
	//RetrieveFeeds函数返回两个值：一组Feed的切片（动态数组），错误值
	//:=运算符用于声明一个变量，同时给这个变量赋予初始值
	feeds, err := RetrieveFeeds()
	if err != nil {
		//Fatal将错误在终端输出，并终止程序
		log.Fatal(err)
	}

	// Create an unbuffered channel to receive match results to display.
	//如果需要声明初始值为零值的变量，应该使用var关键字声明变量，如果提供确切的非零值初始化变量或者使用函数返回值创建变量，应该使用简化变量声明运算符
	results := make(chan *Result)

	// Setup a wait group so we can process all the feeds.
	//如果main函数返回，整个程序也就终止了．GO程序终止时，还会关闭所有之前启动且还在运行的goroutine
	//写并发程序时，最佳做法是，在main函数返回前，清理并终止所有之前启动的goroutine．

	//sync包对WaitGroup跟踪所有启动的goroutine,它是一个计数信号量，我们可以利用它来统计所有的goroutine是不是都完成了工作
	var waitGroup sync.WaitGroup

	// Set the number of goroutines we need to wait for while
	// they process the individual feeds.
	//每个goroutine完成其工作之后，就会递减WaitGroup变量的计数值,当这个值减为0时，我们就知道所有对工作都做完了
	waitGroup.Add(len(feeds))

	// Launch a goroutine for each feed to find the results.
	//for range对feeds切片做迭代，关键字range可以用于迭代数组，字符串，切片，映射和通道
	//使用for range迭代切片时，每次迭代会返回两个值：索引，元素值的一个副本
	//下划线_的作用是占位符，占据类保存range调用返回的索引值的变量的位置．如果要调用的函数返回多个值，而又不需要其中某个值，就可以使用下划线标识符将其忽略．
	for _, feed := range feeds {
		// Retrieve a matcher for the search.
		//从map中查找一个可以用于处理特定数据源类型的数据
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// Launch the goroutine to perform the search.
		//使用go关键字启动一个goroutine，并对这个goroutine做并发调度,
		// 这里启动了一个匿名函数，它有两个参数：一个类型为Matcher，另一个是指向Feed类型值的指针,这意味着feed是一个指针变量．
		//指针变量可以方便地在函数之间共享数据，使用指针变量可以让函数访问并修改一个变量的状态
		//在go中，所有的变量都以值的方式传递
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			//递减WaitGroup的计数
			waitGroup.Done()
			//go语言支持闭包，searchTerm，results，waitGroup都是通过闭包的形式访问
			//因为matcher和feed会随着迭代变化，所以不能通过闭包调用
		}(matcher, feed)
	}

	// Launch a goroutine to monitor when all the work is done.
	go func() {
		// Wait for everything to be processed.
		//阻塞goroutine
		waitGroup.Wait()

		// Close the channel to signal to the Display
		// function that we can exit the program.
		close(results)
	}()

	// Start displaying results as they are available and
	// return after the final result is displayed.
	Display(results)
}

// Register is called to register a matcher for use by the program.
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
