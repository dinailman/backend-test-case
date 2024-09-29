package main

func sumDiagonal(inputMatrix [][]int) (diag1 []int, diag2 []int) {
	j := 0
	k := 0

	for i := 0; i < len(inputMatrix); i++ {
		diag1 = make([]int, len(inputMatrix[i]))
		diag2 = make([]int, len(inputMatrix[i]))
	}
	for i := 0; i < len(inputMatrix); i++ {
		j = i
		diag1[i] = inputMatrix[i][j]

		k = len(inputMatrix[i]) - i - 1
		diag2[i] = inputMatrix[i][k]
	}

	return diag1, diag2
}
