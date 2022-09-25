package main

func main() {
	//i := 1

	//for i <= 3 {
	//	fmt.Println(i)
	//	i++
	//}

	for i := 7; i <= 9; i++ {
		println(i)
	}

	//create loop break and continue

	for {
		println("hello")
		break
	}
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			continue
		}
		println(i)
	}
}
