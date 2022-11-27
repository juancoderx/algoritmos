package main

func countSort(arr []int) []int {
	max := arr[0]

	i := 1
	for i < len(arr) {
		if arr[i] > max {
			max = arr[i]
		}

		i++
	}

	indices := make([]int, max+1)

	j := 0
	for j < len(arr) {
		indices[arr[j]]++

		j++
	}

	k := 1
	for k < len(indices) {
		indices[k] += indices[k-1]

		k++
	}

	result := make([]int, len(arr))

	m := 0
	for m < len(arr) {
		result[indices[arr[m]]-1] = arr[m]
		indices[arr[m]]--

		m++
	}

	return result
}
