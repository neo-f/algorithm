package week21

// https://leetcode-cn.com/problems/walking-robot-simulation/

func robotSim(commands []int, obstacles [][]int) int {
	ans := 0
	// E S W N
	dx := [4]int{1, 0, -1, 0}
	dy := [4]int{0, -1, 0, 1}
	// 初始面向北
	direction := 3
	// x, y
	position := [2]int{0, 0}
	obstaclesSet := make(map[int]struct{})
	calcHash := func(x int, y int) int {
		return (x+30000)*60000 + y + 30000
	}
	for _, o := range obstacles {
		obstaclesSet[calcHash(o[0], o[1])] = struct{}{}
	}

	for _, c := range commands {
		switch c {
		case -1:
			direction = (direction + 1) % 4
		case -2:
			direction = (direction - 1 + 4) % 4
		default:
			for s := 0; s < c; s++ {
				x := position[0] + dx[direction]
				y := position[1] + dy[direction]
				if _, ok := obstaclesSet[calcHash(x, y)]; ok {
					break
				}
				position[0] = x
				position[1] = y
				distance := x*x + y*y
				if distance > ans {
					ans = distance
				}
			}
		}
	}
	return ans
}
