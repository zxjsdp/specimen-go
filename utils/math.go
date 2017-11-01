package utils

// 生成 combinations
//
// Usage:
// utils.GenerateCombinations([]string{"茎", "叶", "花", "果"}, 2, func(c []string) {
//     log.Println(c)
// })
//
// Results:
// [茎 叶]
// [茎 花]
// [茎 果]
// [叶 花]
// [叶 果]
// [花 果]
func GenerateCombinations(source []string, m int, emit func([]string)) {
	s := make([]string, m)
	last := m - 1
	var rc func(int, int)
	rc = func(i, next int) {
		for j := next; j < len(source); j++ {
			s[i] = source[j]
			if i == last {
				emit(s)
			} else {
				rc(i+1, j+1)
			}
		}
		return
	}
	rc(0, 0)
}
