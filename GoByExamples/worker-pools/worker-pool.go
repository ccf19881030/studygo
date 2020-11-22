// Go by Example 中文： 工作池
// http://books.studygolang.com/gobyexample/worker-pools/
// 在这个例子中，我们将看到如何使用Go协程和通道实现一个工作池
package main

import "fmt"
import "time"

// 这是我们将要在多个并发实例中支持的任务了。这些执行者将从 jobs 通道接收任务，并且通过 results 发送对应的结果。我们将让每个任务间隔 1s 来模仿一个耗时的任务。
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "processing job", j)
	time.Sleep(time.Second)
	results <- j * 2
    }
}

func main() {
    startTime := time.Now()
    // 为了使用worker工作池并且收集他们的结果，我们需要2个通道。
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    // 这里启动了3个worker，初始是阻塞的，因为还没有传递任务。
    for w := 1; w <= 3; w++ {
	go worker(w, jobs, results)
    }

    // 这里我们发送9个jobs，然后close这些通道来表示这些就是所有的任务了。
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)

    // 最后，我们收集这些任务的返回值。
    for a := 1; a <= 9; a++ {
        <- results
    }

    endTime := time.Now()
    diffTime := endTime.Sub(startTime).Seconds()
    fmt.Println("任务一共花费的时间为：", diffTime, "秒")
}
