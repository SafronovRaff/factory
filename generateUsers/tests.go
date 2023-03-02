package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var actions = []string{
	"logged in",
	"logged out",
	"create record",
	"delete record",
	"update record",
}

type LogItem struct {
	action    string
	timestamp time.Time
}
type User struct {
	id    int
	email string
	logs  []LogItem
}

func (u User) getActivityInfo() string {
	out := fmt.Sprintf("ID %d | Email: %s \n Activity Log:\n", u.id, u.email)
	for i, items := range u.logs {
		out += fmt.Sprintf("%d. [%s] at %s \n", i+1, items.action, items.timestamp)
	}
	return out
}

func main() {

	rand.Seed(time.Now().Unix())
	users := generateUser(30)
	for _, user := range users {
		seveUserInfo(user)
	}
}

// seveUserInfo - создаёт фал и записывает в него логи пользователей
func seveUserInfo(user User) error {
	fmt.Printf("writing file user ID: %d\n", user.id)

	filename := fmt.Sprintf("logs/uid_%d.txt", user.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = file.WriteString(user.getActivityInfo())
	return err
}

// generateUser - генерирует заданное количество пользователей
func generateUser(count int) []User {
	users := make([]User, count)
	for i := 0; i < count; i++ {
		users[i] = User{
			id:    i + 1,
			email: fmt.Sprintf("user%d@google.com", i+1),
			logs:  generateLogs(rand.Intn(50)), //  рандом до 50 записей
		}
	}
	return users
}

// generateUser - генерирует рандомное количество логов
func generateLogs(count int) []LogItem {
	logs := make([]LogItem, count)
	for i := 0; i < count; i++ {
		logs[i] = LogItem{
			action:    actions[rand.Intn(len(actions)-1)], // рандом от 0 до 5, что бы не выйти за рамки слайса
			timestamp: time.Now(),
		}
	}
	return logs
}
