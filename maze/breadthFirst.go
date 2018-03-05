package maze

func BreadthFirst(m map[int][]int, ent, ext int) []int {
	var p int
	var b []int
	sol := []int{}
	visited := make(map[int]bool)
	notVisited := []int{ent}
	pMap := make(map[int]int)
	pMap[ent] = -1
	for {
		p = notVisited[0]
		b = m[p]
		if p == ext {
			v := pMap[p]
			for v >= 0 {
				sol = append([]int{v}, sol...)
				v = pMap[v]
			}
			return sol

		} else {
			notVisited = append(notVisited[:0], notVisited[1:]...)
			visited[p] = true

			for _, v := range b {
				if !visited[v] {
					pMap[v] = p
					notVisited = append(notVisited, v)
				}
			}
		}

	}
}
