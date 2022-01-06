package main

// practice
// https://github.com/TheAlgorithms/Go/blob/master/sort/mergesort.go
func merge(a []int, b []int) []int {

	var r = make([]int, len(a)+len(b))
	var i = 0
	var j = 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++

		}
	}

	for i < len(a) {
		r[i+j] = a[i]
		i++
	}

	for j < len(b) {
		r[i+j] = b[j]
		j++
	}
	return r
}

func mergesort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	var middle = len(items) / 2
	var a = mergesort(items[:middle])
	var b = mergesort(items[middle:])
	return merge(a, b)

}

func mergeIter(items []int) []int {
	for step := 1; step < len(items); step += step {
		for i := 0; i+step < len(items); i += 2 * step {
			tmp := merge(items[i:i+step], items[i+step:min.Int(i+2*step, len(items))])
			copy(items[i:], tmp)
		}
	}
	return items
}
