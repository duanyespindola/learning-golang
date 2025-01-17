package sum

func Sum(numbers []int) int {
	resp := 0
	for _, value := range numbers {
		resp += value
	}
	return resp
}
