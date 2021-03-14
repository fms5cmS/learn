package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var wg sync.WaitGroup
	length := len(nums)
	workNum := length / 10
	fmt.Println(workNum)
	var ch chan int
	if length%10 == 0 { // 边界条件！
		ch = make(chan int, workNum)
		wg.Add(workNum)
		for i := 0; i < length; i += 10 {
			go sumTime(&wg, nums[i:i+10], ch)
		}
	} else {
		ch = make(chan int, workNum+1)
		wg.Add(workNum + 1)
		for i := 0; i < length; i += 10 {
			if length-i >= 10 {
				go sumTime(&wg, nums[i:i+10], ch)
			} else {
				go sumTime(&wg, nums[i:], ch)
			}
		}
	}
	wg.Wait()
	close(ch)
	sumTime := 0
	for v := range ch {
		sumTime += v
	}
	fmt.Println(sumTime)
}

func sumTime(wg *sync.WaitGroup, data []int, ch chan int) {
	sum := 0
	for _, d := range data {
		sum += d
	}
	ch <- sum
	wg.Done()
}
