#Code optimization with Worker Pool


| workerCount | Time                             |
|-------------|----------------------------------|
| 1           |DONE! Time Elapsed: 111.13 seconds|
| 5           |DONE! Time Elapsed: 30.43 seconds |
| 10          |DONE! Time Elapsed: 20.29 seconds |
| 30          |DONE! Time Elapsed: 14.24 seconds |
| 100         |DONE! Time Elapsed: 11.51 seconds |


````
package main

import (
"fmt"
"log"
"math/rand"
"os"
"time"
)

var actions = []string{"logged in", "logged out", "created record", "deleted record", "updated account"}

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

func main() {
rand.Seed(time.Now().Unix())

	startTime := time.Now()

	users := generateUsers(100)

	for _, user := range users {
		saveUserInfo(user)
	}

	fmt.Printf("DONE! Time Elapsed: %.2f seconds\n", time.Since(startTime).Seconds())
}

func saveUserInfo(user User) {
fmt.Printf("WRITING FILE FOR UID %d\n", user.id)

	filename := fmt.Sprintf("users/uid%d.txt", user.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString(user.getActivityInfo())
	time.Sleep(time.Second)
}

func generateUsers(count int) []User {
users := make([]User, count)

	for i := 0; i < count; i++ {
		users[i] = User{
			id:    i + 1,
			email: fmt.Sprintf("user%d@company.com", i+1),
			logs:  generateLogs(rand.Intn(1000)),
		}
		fmt.Printf("generated user %d\n", i+1)
		time.Sleep(time.Millisecond * 100)
	}

	return users
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
````
