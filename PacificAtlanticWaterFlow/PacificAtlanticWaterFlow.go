package PacificAtlanticWaterFlow

type point struct {
	x int
	y int
}

func pacificAtlantic(matrix [][]int) [][]int {
	dpPacific := make(map[point]bool)
	dpAtlantic := make(map[point]bool)

	var result [][]int
	for i, row := range matrix {
		for j := range row {
			flow(i, j, matrix, dpPacific, make(map[point]struct{}), func(i, j int) bool {
				return i == 0 || j == 0
			})
			flow(i, j, matrix, dpAtlantic, make(map[point]struct{}), func(i, j int) bool {
				return i == len(matrix) - 1 || j == len(matrix[0]) - 1
			})
			if dpPacific[point{i, j}] && dpAtlantic[point{i,j}] {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

func flow(i, j int, matrix [][]int, dp map[point]bool, visited map[point]struct{}, stop func(i, j int) bool) {
	if _, ok := visited[point{i,j}]; ok {
		return
	}

	visited[point{i, j}] = struct{}{}

	if _, ok := dp[point{i, j}]; ok {
		return
	}

	if stop(i, j) {
		dp[point{i, j}] = true
		return
	}

	if i > 0 && matrix[i-1][j] <= matrix[i][j] {
		flow(i-1, j, matrix, dp, visited, stop)
		if dp[point{i-1, j}] {
			dp[point{i, j}] = true
			return
		}
	}

	if i < len(matrix) - 1 && matrix[i+1][j] <= matrix[i][j] {
		flow(i+1, j, matrix, dp, visited, stop)
		if dp[point{i+1, j}] {
			dp[point{i, j}] = true
			return
		}
	}

	if j > 0 && matrix[i][j-1] <= matrix[i][j] {
		flow(i, j-1, matrix, dp, visited, stop)
		if dp[point{i, j-1}] {
			dp[point{i, j}] = true
			return
		}
	}

	if j < len(matrix[0]) - 1 && matrix[i][j+1] <= matrix[i][j] {
		flow(i, j+1, matrix, dp, visited, stop)
		if dp[point{i, j+1}] {
			dp[point{i, j}] = true
			return
		}
	}
}
