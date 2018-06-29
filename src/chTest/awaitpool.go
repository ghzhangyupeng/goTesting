package main

import (
	"time"
	"sync"
	rand2 "math/rand"
	"fmt"
)

type Job struct {
	id int
	randomno int
}

type Result struct {
	job Job
	sumofdigits int
}
// 定义一个job通道
var jobs = make(chan Job, 10)
// 定义一个result通道
var results = make(chan Result, 10)

// 定义一个方法，用于计算结果值，例如：234 -> 9
func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no/=10
	}
	time.Sleep(2* time.Second)
	return sum
}

func worker(wg *sync.WaitGroup)  {
	for job := range jobs {
		output := Result{job,digits(job.randomno)}
		results <- output
	}
	wg.Done()
}
// 创建n个worker的连接池
func createWorkerPool(noOfWorkers int)  {
	var wg sync.WaitGroup
	for i := 0;i < noOfWorkers ; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}
// 模拟创建n个任务工作
func allocate(noOfJobs int) {
	for i:=0; i < noOfJobs; i++ {
		randomno := rand2.Intn(999)
		job :=Job{i, randomno}
		// 生成任务放入任务信道中
		jobs <- job
	}
	// 只是告诉信道数据已经发送完了， 如果没有close，后面其他协程执行时 range 信道ch， 会出现死锁问题，因为没有再往信道里面放数据。
	close(jobs)
}

func result(done chan bool)  {
	for result := range results{
		fmt.Printf("Job id %d, input random no %d, sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
	fmt.Printf("blocking ===== ")
}

func main()  {
	startTime := time.Now()
	noOfJobs := 20
	// 模拟创建100个任务
	go allocate(noOfJobs)
	// 用于堵塞标记
	done := make(chan bool, 2)
	// 对一个100个任务进行结果计算
	go result(done)
	noOfWorkers := 10
	// 创建10个worker的工作池
	createWorkerPool(noOfWorkers)
	//fmt.Println("test ing")

	fmt.Println(<-done)
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken", diff.Seconds(), "seconds")
}






