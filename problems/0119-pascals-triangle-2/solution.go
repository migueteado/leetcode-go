package p0119

func pascalsTriangle2(rowIndex int) []int {
	var prev []int
	for i := 0; i <= rowIndex; i++ {
		var curr []int
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				curr = append(curr, 1)
			} else {
				curr = append(curr, prev[j-1]+prev[j])
			}
		}
		prev = curr
	}
	return prev
}
