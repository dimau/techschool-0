package main

import (
	"fmt"
	"math/rand"
	"time"

	"techschool-0/workerpool"
)

func main() {
	// Create a new Worker Pool
	pool := workerpool.NewPool(5)

	// Create a lot of Tasks
	go func() {
		for {
			taskID := rand.Intn(100) + 20

			if taskID%7 == 0 {
				pool.Stop()
			}

			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			task := workerpool.NewTask(func(data interface{}) error {
				taskID := data.(int)
				time.Sleep(100 * time.Millisecond)
				fmt.Printf("Task %d processed\n", taskID)
				return nil
			}, taskID)
			pool.AddTask(task)
		}
	}()

	// Run our Worker Pool
	pool.RunBackground()
}
