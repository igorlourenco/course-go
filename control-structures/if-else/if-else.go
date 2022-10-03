package main

func getGradeSituation(grade float64) {
	if grade >= 6 {
		println("You are approved with", grade)
	} else {
		println("You are disapproved with", grade)
	}
}

func main() {
	getGradeSituation(6.2)
	getGradeSituation(5.8)
}
