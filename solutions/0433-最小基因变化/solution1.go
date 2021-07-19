package main

func minMutation(start string, end string, bank []string) int {
	// 基因库集合
	bankSet := make(map[string]struct{}, len(bank))
	for _, b := range bank {
		bankSet[b] = struct{}{}
	}
	genes := [4]string{"A", "C", "G", "T"}
	// 记录层级
	depth := make(map[string]int)

	// 广搜队列
	queue := []string{start}
	for len(queue) != 0 {
		// Pop 队头
		str := queue[0]
		queue = queue[1:]

		// 遍历每个元素
		for i := 0; i < 8; i++ {
			// 遍历每种基因
			for j := 0; j < 4; j++ {
				if string(str[i]) == genes[j] {
					continue
				}
				// 变成的新基因串
				ns := str[:i] + string(genes[j]) + str[i+1:]
				// 如果基因串不合法，跳过
				if _, ok := bankSet[ns]; !ok {
					continue
				}
				if _, ok := depth[ns]; !ok {
					depth[ns] = depth[str] + 1
					if ns == end {
						return depth[ns]
					}
					queue = append(queue, ns)
				}
			}
		}
	}
	return -1
}
