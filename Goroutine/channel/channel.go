package channel

import (
	"fmt"
	"sync"
	"time"
)

func Channel() {
	var wg sync.WaitGroup

	// channel không buffer bắt buộc đủ 2 go routine 1 cái gửi 1 cái nhận
	// bản thân main cũng là 1 go routine , mỗi lần gửi nhận chỉ 1 giá trị xong lần gửi nhận này mới tới lân kế
	// Ai tới trước thì phải đứng chờ người kia:
	// - Sender gửi trước  -> sender bị block, chờ tới khi receiver nhận -> cả 2 chạy tiếp
	// - Receiver nhận trước -> receiver bị block, chờ tới khi sender gửi -> cả 2 chạy tiếp

	channel := make(chan int)

	//sender
	wg.Go(func() {
		channel <- 1
		channel <- 2
		channel <- 3
		channel <- 4 // deadlock ở đây go routine này ko được kết thúc nếu bên dưới không nhận
	})

	//receiver
	wg.Go(func() {
		fmt.Println(<-channel)
	})

	
	fmt.Println(<-channel)

	time.Sleep(1*time.Second)

	wg.Go(func() {
		fmt.Println(<-channel)
	})

	fmt.Println(<-channel)
	wg.Wait()


	channelB := make(chan string)

	go func ()  {
		defer close(channelB)
		channelB <- "HAHA"
		channelB <- "HIHI"
		
	}()

	for value := range channelB{
		fmt.Println(value)
	}

	fmt.Println("Kết Thúc")
}
