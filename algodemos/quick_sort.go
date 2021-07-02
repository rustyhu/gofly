package algodemos

// ref https://github.com/TheAlgorithms/Go/tree/master/sorts

func partition(a []int) int {
	r := len(a) - 1
	midNumber := a[r]

	i := 0
	for j := 0; j < r; j++ {
		if a[j] < midNumber {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[r] = a[r], a[i]
	return i
}

func Qsort(a []int) {
	if len(a) <= 1 {
		return
	}

	midx := partition(a)
	Qsort(a[:midx])
	Qsort(a[midx+1:])
}
