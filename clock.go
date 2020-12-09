package main

import "fmt"
// import "time"

func main() {
	
	// now := time.Now().Unix()
	// var now int64
	// save := time.Now().Unix()
	// delete := save.Add(3 * time.Second)
	// fmt.Println(save)
	// fmt.Println(delete)
	// yes := now.Equal(delete)
	timeStamp := make([]int64, 10)
	timeStamp[2:7] = 2

	fmt.Println(timeStamp)

	// for {
	// 	now = time.Now().Unix()
	// 	if now - save >= 3 {
	// 		save = time.Now().Unix()
	// 		fmt.Println(time.Now())
	// 	}

		// time.Sleep(1 * time.Second)
	// }
}