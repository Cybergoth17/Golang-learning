package main

import "fmt"

var message string

type User struct {
	name   string
	age    int
	sex    string
	weight int
	height int
}

func (u User) receiver() {
	fmt.Println(u.name, u.age, u.sex, u.weight, u.height)
}
func init() {
	message = "I am using init() fucntion right now..."
}
func main() {
	user1 := User{"Lana", 17, "female", 48, 163}
	user1.receiver()

	users := map[string]int{
		"Anel":    18,
		"Aruzhan": 18,
		"Aliya":   16,
	}
	// add new key and value to users map
	users["Hamza"] = 17
	// delete key from users map
	delete(users, "Aruzhan")
	for key, value := range users {
		fmt.Println(key, value)
	}
	// from init function
	fmt.Println(message)
	//find minimum
	fmt.Println(FindMin(1, 25, 589, 4784, -2, 887, 55888, 4))

	//matrix output
	a := [5][2]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}}

	for i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("a[%d][%d]=%d\n", i, j, a[i][j])
		}
	}
	//anonimus function
	func() {
		fmt.Println("you are gay")
	}()
	inc := increment()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	//pointers
	message1 := "I will be a good backender"
	fmt.Println(message1)
	changeMessage(&message1)
	fmt.Println(message1)

	list := []int{1, 2, 4, 8, 16}
	fmt.Println(double(list))
}
func double(nums []int) []int {
	res := make([]int, 0, len(nums))
	for _, num := range nums {
		res = append(res, num*2)
	}
	return res
}

//pointers function
func changeMessage(message1 *string) {
	*message1 += " {from changeMessage()}"
}

//increment
func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func FindMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	min := numbers[0]
	for _, i := range numbers {
		if i < min {
			min = i
		}
	}
	return min
}
