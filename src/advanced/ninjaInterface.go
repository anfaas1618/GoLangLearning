package main

type NinjaStar struct {
	owner string
}

func (NinjaStar) attack() {
	println("throwing nija star")
}

type NinjaSword struct {
	owner string
}

func (n NinjaSword) attack() {
	println("attacking sward")
}

type wepons interface {
	attack()
}

func main() {
	ninjas := []NinjaStar{{"owner"}, {"second"}}
	for _, ninja := range ninjas {
		ninja.attack()
	}
	var weapons = []wepons{NinjaSword{"yes"}, NinjaStar{"yla"}}

	for _, value := range weapons {
		value.attack()
	}
}
