#### goroutine练习笔记

> ##### 1.模拟乒乓球
>
> ```go
> package main
> 
> import (
> 	"fmt"
> 	"time"
> )
> 
> type Ball struct {
> 	hits int
> }
> 
> func main() {
> 	table := make(chan *Ball)
> 	go player("ping", table)
> 	go player("pong", table)
> 
> 	table <- new(Ball)
> 	time.Sleep(time.Second)
> 	<- table
> 	defer close(table)
> }
> 
> func player(name string, table chan *Ball) {
> 	for {
> 		ball, ok := <- table
> 		if !ok {
> 			break
> 		}
> 		ball.hits += 1
> 		fmt.Println(ball.hits, name)
> 		table <- ball
> 		time.Sleep(time.Millisecond)
> 	}
> 	return
> }
> ```
>
> ##### 2.打印12AB34CD……2526YZ
>
> ```go
> package main
> 
> import (
> 	"fmt"
> 	"time"
> )
> 
> func main() {
> 	chanNum := make(chan bool)
> 	chanByte := make(chan bool)
> 	go PrintNum(chanNum, chanByte)
> 	go PrintByte(chanNum, chanByte)
> 
> 	chanNum <- true
> 	time.Sleep(time.Second)
> 	defer close(chanNum)
> 	defer close(chanByte)
> }
> 
> func PrintNum(chanNum, chanByte chan bool) {
> 	for i := 1; i <= 26; i += 2 {
> 		<- chanNum
> 		fmt.Printf("%d", i)
> 		fmt.Printf("%d", i+1)
> 		chanByte <- true
> 		time.Sleep(time.Millisecond)
> 	}
> 	return
> }
> 
> func PrintByte(chanNum, chanByte chan bool) {
> 	for i := 'A'; i <= 'Z'; i += 2{
> 		<- chanByte
> 		fmt.Printf("%c", i)
> 		fmt.Printf("%c", i+1)
> 		chanNum <- true
> 		time.Sleep(time.Millisecond)
> 	}
> 	return
> }
> ```
>
> ##### 3.cat dog mouse循环打印
>
> ```go
> package main
> 
> import (
> 	"fmt"
> 	"time"
> )
> 
> func main() {
> 	chanCat := make(chan *int)
> 	chanDog := make(chan *int)
> 	chanMouse := make(chan *int)
> 
> 	go Cat(chanMouse, chanCat)
> 	go Dog(chanCat, chanDog)
> 	go Mouse(chanDog, chanMouse)
> 
> 	var times = 0
> 	chanMouse <- &times
> 	time.Sleep(time.Second)
> 
> 	defer close(chanMouse)
> 	defer close(chanDog)
> 	defer close(chanCat)
> 
> }
> 
> func Cat (chanMouse, chanCat chan *int) {
> 	for i := 0; i < 100; i++ {
> 		times := <- chanMouse
> 		*times += 1
> 		fmt.Println(*times, "Cat")
> 		chanCat <- times
> 		time.Sleep(time.Millisecond)
> 	}
> 	return
> }
> 
> func Dog (chanCat, chanDog chan *int) {
> 	for i := 0; i < 100; i++ {
> 		times := <- chanCat
> 		*times += 1
> 		fmt.Println(*times, "Dog")
> 		chanDog <- times
> 		time.Sleep(time.Millisecond)
> 	}
> 	return
> }
> func Mouse (chanDog, chanMouse chan *int) {
> 	for i := 0; i < 100; i++ {
> 		times := <- chanDog
> 		*times += 1
> 		fmt.Println(*times, "Mouse")
> 		chanMouse <- times
> 		time.Sleep(time.Millisecond)
> 	}
> 	return
> }
> 
> //使用sync.WaitGroup
> package main
> 
> import (
> 	"fmt"
> 	"sync"
> )
> 
> //var wg sync.WaitGroup
> 
> func main() {
> 	wg := sync.WaitGroup{}
> 	//注意使用有缓冲channel
> 	//无缓冲通道要求发送goroutine和接收goroutine同时准备好，才能完成和发送接收操作
> 	//如果两个goroutine没有同时准备好，通道会导致先执行或者接收操作的goroutine阻塞等待。
> 	chanCat := make(chan *int, 1)
> 	chanDog := make(chan *int, 1)
> 	chanMouse := make(chan *int, 1)
> 
> 	wg.Add(3)
> 	go Cat(chanMouse, chanCat, &wg)
> 	go Dog(chanCat, chanDog, &wg)
> 	go Mouse(chanDog, chanMouse, &wg)
> 
> 	var times = 0
> 	chanMouse <- &times
> 
> 	wg.Wait()
> 	defer close(chanCat)
> 	defer close(chanDog)
> 	defer close(chanMouse)
> }
> 
> func Cat (chanMouse, chanCat chan *int, wg *sync.WaitGroup) {
> 	defer wg.Done()
> 	for i := 0; i < 100; i++ {
> 		times := <- chanMouse
> 		*times += 1
> 		fmt.Println(*times, "Cat")
> 		chanCat <- times
> 	}
> }
> 
> func Dog (chanCat, chanDog chan *int, wg *sync.WaitGroup) {
> 	defer wg.Done()
> 	for i := 0; i < 100; i++ {
> 		times := <- chanCat
> 		*times += 1
> 		fmt.Println(*times, "Dog")
> 		chanDog <- times
> 	}
> }
> func Mouse (chanDog, chanMouse chan *int, wg *sync.WaitGroup) {
> 	defer wg.Done()
> 	for i := 0; i < 100; i++ {
> 		times := <- chanDog
> 		*times += 1
> 		fmt.Println(*times, "Mouse")
> 		chanMouse <- times
> 	}
> }
> ```
>
> ##### 4.使用channel限制goroutine数量
>
> ```go
> package main
> 
> import (
> 	"fmt"
> 	"runtime"
> 	"sync"
> 	"time"
> )
> 
> func main() {
> 	wg := sync.WaitGroup{}
> 	taskInt := 10
> 	ch := make(chan bool, 3)
> 
> 	for i := 0; i <= taskInt; i++ {
> 		wg.Add(1)
> 		ch <- true
> 		go worker(ch, &wg, i)
> 		time.Sleep(time.Second)
> 	}
> 	wg.Wait()
> 	defer close(ch)
> }
> 
> func worker(ch chan bool, wg *sync.WaitGroup, i int) {
> 	defer wg.Done()
> 	fmt.Println("go func: ", i, "goroutine number: ", runtime.NumGoroutine())
> 	time.Sleep(5*time.Second)
> 	<- ch
> 	return
> }
> ```
>
> 