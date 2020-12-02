package Week02

import "fmt"

func main() {
	var id int = 1
	person ,err := GetPersonById(id)
	if err != nil {
		// 记录日志，返回给上层错误信息
		fmt.Println("")
	}

	fmt.Println(person)
}