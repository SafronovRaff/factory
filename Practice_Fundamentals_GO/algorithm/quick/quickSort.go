package quick

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}

	quickSort(ar)

	fmt.Println(ar)
}

func quickSort(ar []int) {
	if len(ar) < 2 {
		return
	}
	left, right := 0, len(ar)-1
	povorIndex := rand.Int() % len(ar)
	ar[povorIndex], ar[right] = ar[right], ar[povorIndex]
	for i := 0; i < len(ar); i++ {
		if ar[i] < ar[right] {
			ar[i], ar[left] = ar[left], ar[i]
			left++
		}
	}
	ar[left], ar[right] = ar[right], ar[left]
	quickSort(ar[:left])
	quickSort(ar[left+1:])
	return
}
