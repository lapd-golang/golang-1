package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id int
	randomno int
}

type Result struct{
	job Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func sum_digits(number int)int {
	sum := 0
	for number !=0 {
		i:= number%10
		sum += i
		number = number/10
	}
	time.Sleep(2*time.Second)
	return sum
}
func worker(wg *sync.WaitGroup){
	for job := range jobs{
		output := Result{job, sum_digits(job.randomno)}
		results <- output
	}

	wg.Done()
}

func create_worker_pool(num_workers int){
	var wg sync.WaitGroup
	for i:=0; i < num_workers; i++{
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(num_jobs int){
	for i := 0; i < num_jobs; i++{
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}
func result(done chan bool){
	for result := range results{
		fmt.Printf("Job id %d, input random no %d, sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func main(){
	startTime := time.Now()
	num_jobs := 100
	go allocate(num_jobs)
	done := make(chan bool)
	go result(done)
	num_workers := 10
	create_worker_pool(num_workers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

