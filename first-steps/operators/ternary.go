package main

// there is no ternary operator in Go
func getGradeSituation(grade int) string {
	if grade >= 6 {
		return "approved"
	}
	return "disapproved"
}

func main() {
	grade := 6
	println("You are", getGradeSituation(grade))
}
