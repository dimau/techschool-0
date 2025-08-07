package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"techschool-0/workerpool"
)

func main() {
	// Вешаем нашу функцию–обработчик на все запросы по адресу "/hello"
	http.HandleFunc("/hello", helloHandler)

	// Запускаем веб–сервер на локальной машине на порту 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// Если не удалось запустить веб–сервер, то выдаем ошибку
		panic(err)
	}

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

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	randomNumber := rand.Intn(100)
	_, err := writer.Write([]byte("Hello, world! The random number is " + strconv.Itoa(randomNumber)))
	if err != nil {
		panic(err)
	}
}
