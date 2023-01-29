package main

import "fmt"

func main() {

fmt.Println("aca empieza el primer for loop")
i := 1
for i <= 30 {
	fmt.Println(i)
	i = i + 1
}

fmt.Println("aca empieza el segundo for loop")
for j := 7; j <= 9; j++ {
	fmt.Println(j)
}

for {
	fmt.Print("this is a loop ")
	break
}

for n := 0; n <= 5; n++ {
	if n%2 == 0 {
		continue
	}
	fmt.Println(n)
}
}