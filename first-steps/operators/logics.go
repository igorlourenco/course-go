package main

func shopping(workAtTuesday, workAtWednesday bool) (bool, bool, bool) {
	tvL := workAtTuesday && workAtWednesday
	tvS := workAtTuesday != workAtWednesday // exclusive or
	iceCream := workAtTuesday || workAtWednesday

	return tvL, tvS, iceCream
}

func main() {
	println(shopping(false, false))
}
