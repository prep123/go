# channel

一般channel的声明形式为：
var chanName chan ElementType
与一般的变量声明不同的地方仅仅是在类型之前加了chan关键字。ElementType指定这个
channel所能传递的元素类型。举个例子，我们声明一个传递类型为int的channel：
var ch chan int
或者，我们声明一个map，元素是bool型的channel:
var m map[string] chan bool
上面的语句都是合法的。
定义一个channel也很简单，直接使用内置的函数make()即可：
ch := make(chan int)
这就声明并初始化了一个int型的名为ch的channel。
在channel的用法中，最常见的包括写入和读出。将一个数据写入（发送）至channel的语法
很直观，如下：
ch <- value
向channel写入数据通常会导致程序阻塞，直到有其他goroutine从这个channel中读取数据。从
channel中读取数据的语法是
value := <-ch
如果channel之前没有写入数据，那么从channel中读取数据也会导致程序阻塞，直到channel
中被写入数据为止。我们之后还会提到如何控制channel只接受写或者只允许读取，即单向
channel。