package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type T struct {
	name  string
	value int
}

func human(name string, age int) string {
	// age는 int 타입이므로 문자열과 덧붙이기 위해서는 문자열로 변환
	// Sprintf 함수는 문자열 형식을 지정하여 값을 포맷팅한 뒤 문자열로 반환하는 기능을 제공
	str := fmt.Sprintf("Hello, My name is %s and I'm %d years old.", name, age)
	return str
}

func main() {
	fmt.Println("Hello, World!")

	result := human("Alice", 25)
	fmt.Println(result)

	strArray := [...]string{"Kim", "Lee", "Park"}

	for _, str := range strArray {
		fmt.Println(str)
	}

	//gin.SetMode(gin.ReleaseMode) 릴리즈 모드 활성화 하고 싶으면 주석 해제
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": result,
			"array":   strArray,
		})
	})
	r.Run(":3000") // port

}
