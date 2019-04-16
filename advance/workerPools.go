/*
	这里使用goroutine和channel实现一个worker pools
*/
package main

import "log"
import "time"

//Here’s the worker, of which we’ll run several concurrent instances.
//These workers will receive work on the jobs channel and
//send the corresponding results on results. We’ll sleep a second per job to simulate an expensive task.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		log.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		log.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	//init log,打印文件名，行号
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	//In order to use our pool of workers we need to send them work
	//and collect their results. We make 2 channels for this.
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	//This starts up 3 workers, initially blocked because there are no jobs yet.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	//Here we send 5 jobs and then close that channel to indicate that’s all the work we have.
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	//Finally we collect all the results of the work.
	for a := 1; a <= 5; a++ {
		// <-results
		result := <-results
		log.Println("the ", a, " result= ", result)
	}
}

/*
Our running program shows the 5 jobs being executed by various workers.
尽管这个工作量是5秒，但是我们的程序却只花费了2秒，因为有三个worker在并行的执行
运行： go run workerPools.go
结果：
2017/11/10 15:50:14 workerPools.go:13: worker 3 started  job 1
2017/11/10 15:50:14 workerPools.go:13: worker 2 started  job 3
2017/11/10 15:50:14 workerPools.go:13: worker 1 started  job 2
2017/11/10 15:50:15 workerPools.go:15: worker 1 finished job 2
2017/11/10 15:50:15 workerPools.go:15: worker 3 finished job 1
2017/11/10 15:50:15 workerPools.go:13: worker 3 started  job 5
2017/11/10 15:50:15 workerPools.go:13: worker 1 started  job 4
2017/11/10 15:50:15 workerPools.go:42: the  1  result=  4
2017/11/10 15:50:15 workerPools.go:42: the  2  result=  2
2017/11/10 15:50:15 workerPools.go:15: worker 2 finished job 3
2017/11/10 15:50:15 workerPools.go:42: the  3  result=  6
2017/11/10 15:50:16 workerPools.go:15: worker 3 finished job 5
2017/11/10 15:50:16 workerPools.go:42: the  4  result=  10
2017/11/10 15:50:16 workerPools.go:15: worker 1 finished job 4
2017/11/10 15:50:16 workerPools.go:42: the  5  result=  8

*/
