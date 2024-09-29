package main

func queryEachInput(strInput []string, strQuery []string) (resultInput []int) {

	var countInput = make([]int, len(strQuery))

	for _, v := range strInput {
		for index, q := range strQuery {
			if v == q {
				countInput[index] = countInput[index] + 1
			}
		}
	}
	resultInput = countInput

	return resultInput
}
