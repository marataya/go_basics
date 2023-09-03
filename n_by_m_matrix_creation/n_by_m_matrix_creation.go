package n_by_m_matrix_creation

import "fmt"

func createMatrix(rows, cols int) [][]int {
	result := [][]int{}
	for i := 0; i < rows; i++ {
		row := make([]int, 0)
		for j := 0; j < cols; j++ {
			row = append(row, (i+1)*(j+1))
		}
		result = append(result, row)
	}

	return result
}

func test1(rows, cols int) {
	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
	matrix := createMatrix(rows, cols)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("======= END REPORT =======")
}

func main() {
	test1(5, 4)
	test1(10, 10)
	test1(100, 100)
}
