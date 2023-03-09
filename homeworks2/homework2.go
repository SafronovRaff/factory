package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

type logItem struct {
	action    string
	timestamp time.Time
}

type User struct {
	id    int
	email string
	logs  []logItem
}

func (u User) getActivityInfo() string {
	output := fmt.Sprintf("UID: %d; Email: %s;\nActivity Log:\n", u.id, u.email)
	for index, item := range u.logs {
		output += fmt.Sprintf("%d. [%s] at %s\n", index, item.action, item.timestamp.Format(time.RFC3339))
	}

	return output
}

var (
	workerCount, result = 100, 100
	actions             = []string{"logged in", "logged out", "created record", "deleted record", "updated account"}
	startTime           = time.Now()
	wg                  sync.WaitGroup
)

func init() {
	rand.Seed(time.Now().Unix())
}
func main() {
	jobs := make(chan int, result)

	taskQueue := make(chan User, result) // канал для очереди задач

	generateUsers(result, jobs, taskQueue) // генерируем 100 пользователей

	generationRes(result, jobs, &wg)

	saveUserInfo(workerCount, taskQueue, &wg)

	wg.Wait() // блокируем до тех пор, пока не выполним все горутины

	fmt.Printf("DONE! Time Elapsed: %.2f seconds\n", time.Since(startTime).Seconds())
}

func generationRes(result int, jobs chan int, wg *sync.WaitGroup) {
	for i := 0; i < result; i++ {
		wg.Add(1) // добавляем 1 задачу
		jobs <- i
	}
}

func saveUserInfo(workerCount int, taskQueue chan User, wg *sync.WaitGroup) {
	for i := 0; i < workerCount; i++ {
		go func() {
			for user := range taskQueue {
				fmt.Printf("WRITING FILE FOR UID %d\n", user.id)
				filename := fmt.Sprintf("users/uid_%d.txt", user.id)
				file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
				if err != nil {
					log.Fatal(err)
				}

				file.WriteString(user.getActivityInfo())

				time.Sleep(time.Second)
				wg.Done()
			}
		}()
	}
}

func generateUsers(count int, jobs chan int, taskQueue chan User) {

	for i := 0; i < count; i++ {

		go func() {
			for i := range jobs {
				taskQueue <- User{
					id:    i + 1,
					email: fmt.Sprintf("user%d@company.com", i+1),
					logs:  generateLogs(rand.Intn(1000)),
				}
				fmt.Printf("generated user %d\n", i+1)
				time.Sleep(time.Millisecond * 100)
			}
			close(taskQueue)
		}()
	}
}
func generateLogs(count int) []logItem {
	logs := make([]logItem, count)

	for i := 0; i < count; i++ {
		logs[i] = logItem{
			action:    actions[rand.Intn(len(actions)-1)],
			timestamp: time.Now(),
		}
	}

	return logs
}
