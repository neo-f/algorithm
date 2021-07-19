package week12

import . "algorithm-go/utils"

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	area := 0
	for i < j {
		// 面积由短的一根的高度决定
		h := MinInt(height[i], height[j])
		area = MaxInt(area, h*(j-i))
		// 比较两边的柱子，如果某侧的柱子比较矮，因为面积是由矮柱子决定的，所以较矮柱子一侧的面积在往中间方向走的时候已经是最大的了，
		// 需要该侧柱子往中间移动，检查是否可能有更高的柱子能组成更大的面积
		// 一样高的时候就一起移动，或者移动随意一侧继续计算即可
		delta := height[i] - height[j]
		switch {
		case delta > 0:
			j--
		case delta < 0:
			i++
		default:
			j--
			i++
		}
	}
	return area
}
