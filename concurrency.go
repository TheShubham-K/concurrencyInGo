package main

type IBankAccount interface {
	GetBalance() int // 100 = 1 dollar
	Deposit(amount int)
	Withdraw(amount int) error
}

func main() {

	// calling two goroutine
	// go count("sheep")
	// count("fish")

	// go count("fish")
	// go count("sheep")

	// // it's way to interrupt the go routine
	// // A Blocking call
	// fmt.Scanln()

	// sync and wait call
	// var wg sync.WaitGroup
	// wg.Add(1)

	// go func() {
	// 	count("fish")
	// 	count("sheep")
	// 	wg.Done()
	// }()
	// wg.Wait()

	// passing channels to goroutine
	// c := make(chan string)
	// go count("sheep", c)
	// go count("fish", c)

	// // does the same thing as that of the condition check
	// // inside the for loop jst synatics sugar on it with for loop

	// for msg := range c {

	// 	// open checks for the completion of the iterations
	// 	// if its complete it terminates the call and return
	// 	// without any deadlock

	// 	// msg, open := <-c
	// 	// if !open {
	// 	// 	break
	// 	// }
	// 	fmt.Println(msg)
	// }

	// Specifing The Capacity
	// A Buffered channel with capacity defined
	// c := make(chan string, 2)
	// c <- "hello"

	// msg := <-c
	// fmt.Println(msg)

	// c1 := make(chan string)
	// c2 := make(chan string)

	// go func() {
	// 	for {
	// 		c1 <- "Every 500ms"
	// 		time.Sleep(time.Millisecond * 500)
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		c2 <- "Every 2000ms"
	// 		time.Sleep(time.Second * 2)
	// 	}
	// }()

	// for {
	// 	// this way is very slow so we will use select statement instead
	// 	// fmt.Println(<-c1)
	// 	// fmt.Println(<-c2)
	// 	// more effiecent way to make channels receieve and
	// 	// display only those channel which are available
	// 	select {
	// 	case msg1 := <-c1:
	// 		fmt.Println(msg1)
	// 	case msg2 := <-c2:
	// 		fmt.Println(msg2)
	// 	}

	// }

	// worker and loader concept

	// jobs := make(chan int, 100)
	// results := make(chan int, 100)

	// go worker(jobs, results)

	// for i := 0; i < 100; i++ {
	// 	jobs <- i
	// }

	// close(jobs)

	// for j := 0; j < 100; j++ {
	// 	fmt.Println(<-results)
	// }

	// worker example
	// dataChan := make(chan int)
	// go func() {
	// 	wg := sync.WaitGroup{}
	// 	for i := 0; i < 1000; i++ {
	// 		wg.Add(1)
	// 		go func() {
	// 			defer wg.Done()
	// 			results := DoWork()
	// 			dataChan <- results
	// 		}()
	// 	}
	// 	wg.Wait()
	// 	close(dataChan)
	// }()

	// for n := range dataChan {
	// 	fmt.Printf("n = %d\n", n)
	// }

	// clearTimeout()
	// timeoutContext, cancel := context.WithTimeout(context.Background(),
	// 	time.Millisecond*1500)

	// defer cancel()

	// // create HTTP request
	// req, err := http.NewRequestWithContext(timeoutContext,
	// 	http.MethodGet,
	// 	"http://placeholder.it/2000x2000",
	// 	nil)

	// if err != nil {
	// 	panic(err)
	// }

	// // perform HTTP request
	// res, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	panic(err)
	// }

	// defer res.Body.Close()

	// // get data from http response
	// imageData, err := ioutil.ReadAll(res.Body)
	// fmt.Printf("Downloaded image of size %d\n", len(imageData))

	// r := gin.Default()

	// r.GET("/home", func(ctx *gin.Context) {

	// 	timeoutContext, cancel := context.WithTimeout(ctx.Request.Context(), time.Second*5)
	// 	defer cancel()

	// 	// create HTTP request
	// 	req, err := http.NewRequestWithContext(timeoutContext,
	// 		http.MethodGet,
	// 		"http://yahoo.com",
	// 		nil)

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	// perform HTTP request
	// 	res, err := http.DefaultClient.Do(req)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	defer res.Body.Close()

	// 	// get data from http response
	// 	data, err := ioutil.ReadAll(res.Body)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	ctx.Data(200, "text/html", data)
	// })
	// r.Run("0.0.0.0:8080")

	// wf := NewWellsFargo()

	// wf.Deposit(200)
	// wf.Deposit(500)
	// currentBalance := wf.GetBalance()
	// fmt.Printf("WF Balance: %d\n", currentBalance)

	// myAccounts := []IBankAccount{
	// 	NewWellsFargo(),
	// 	NewBitcoinAccount(),
	// }

	// for _, account := range myAccounts {
	// 	account.Deposit(500)
	// 	if err := account.Withdraw(70); err != nil {
	// 		fmt.Printf("ERR: %d\n", err)
	// 	}
	// 	balance := account.GetBalance()
	// 	fmt.Printf("balance = %d\n", balance)
	// }

}

// func DoWork() int {
// 	time.Sleep(time.Second)
// 	return rand.Intn(100)
// }

// func worker(jobs <-chan int, results chan<- int) {
// 	for n := range jobs {
// 		results <- fib(n)
// 	}
// }

// func fib(n int) int {
// 	if n <= 1 {
// 		return n
// 	}
// 	return fib(n-1) + fib(n-2)
// }

// func count(thing string) {

// 	for i := 1; true; i++ {
// 		fmt.Println(i, thing)
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }

// func count(thing string) {

// 	for i := 1; i <= 5; i++ {
// 		fmt.Println(i, thing)
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }

// func count(thing string, c chan string) {

// 	for i := 1; i <= 5; i++ {
// 		c <- thing
// 		time.Sleep(time.Millisecond * 500)
// 	}
// 	close(c)
// }
