package waitgroup

import (
	"fmt"
	"sync"
	"time"
)

func Task(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Bắt đầu task %d\n", num)
	time.Sleep(1 * time.Second)
	fmt.Printf("Kết thúc task %d\n", num)
}

func TaskAutoAddDone(num int) {
	fmt.Printf("Bắt đầu task %d\n", num)
	time.Sleep(1 * time.Second)
	fmt.Printf("Kết thúc task %d\n", num)
}

func WaitGroup() {
	start := time.Now()
	//khai báo wg hoặc var wg sync.WaitGroup -- khai báo bộ đếm đến 1 thì nhớ done nhưng thôi sài wg.GO(fun(){})
	wg := sync.WaitGroup{}

	wg.Add(1)
	go Task(1, &wg)
	wg.Add(1)
	go Task(2, &wg)
	wg.Add(1)
	go Task(3, &wg)
	wg.Add(1)
	go Task(4, &wg)

	wg.Go(func() {
		TaskAutoAddDone(5)
	})

	wg.Go(func() {
		TaskAutoAddDone(6)
	})

	wg.Wait()

	//HÀM GO LUÔN KẾT THÚC KHI MAIN KẾT THÚC BẤT CHẤP GOROUTINE CHƯA XONG
	//time.Sleep(2 * time.Second)

	fmt.Printf("Tổng thời gian thực thi: %v", time.Since(start))
}
