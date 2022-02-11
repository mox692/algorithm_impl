package heap

// len, arr
func heap(n int, input []int) []int {
	// heapを構築
	var heap []int
	for i := 0; i < n; i++ {
		heap = appendHeap(heap, input[i])
	}
	return heap
}

func appendHeap(heap []int, appended int) []int {
	targ := len(heap)
	heap = append(heap, appended)
	for {
		parent := (targ - 1) >> 1
		if targ == 0 {
			break
		}
		if heap[targ] > heap[parent] {
			heap = swap(heap, targ, parent)
			targ = parent
			continue
		}
		break
	}
	return heap
}

// TODO: impl
func deleteTop(heap []int) []int {
	return nil
}

func swap(arr []int, i, j int) []int {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
	return arr
}

func Map[T any, K any](from []T, mapFucn func([]T) []K) []K {
	return mapFucn(from)
}
