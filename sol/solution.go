package sol

type Pair struct {
	row, col int
}

func orangesRotting(grid [][]int) int {
	ROW := len(grid)
	COL := len(grid[0])
	queue := []Pair{}
	fresh := 0
	time := 0
	// collect rotten orange, fresh number
	for row := 0; row < ROW; row++ {
		for col := 0; col < COL; col++ {
			orange := grid[row][col]
			if orange == 2 {
				queue = append(queue, Pair{row: row, col: col})
			}
			if orange == 1 {
				fresh += 1
			}
		}
	}
	// infect directions
	directions := []Pair{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
	// bfs infect
	for len(queue) > 0 && fresh != 0 {
		qLen := len(queue)
		for idx := 0; idx < qLen; idx++ {
			top := queue[0]
			queue = queue[1:]
			for _, direction := range directions {
				shift_row := top.row + direction.row
				shift_col := top.col + direction.col
				if shift_row < 0 || shift_row >= ROW || shift_col < 0 || shift_col >= COL ||
					grid[shift_row][shift_col] != 1 {
					continue
				}
				queue = append(queue, Pair{row: shift_row, col: shift_col})
				grid[shift_row][shift_col] = 2
				fresh -= 1
			}
		}
		time++
	}
	if fresh != 0 {
		return -1
	}
	return time
}
