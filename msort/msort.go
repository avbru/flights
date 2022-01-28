package msort

import "sync"

func MergeSort(arr []int) (res []int) {
	var wg sync.WaitGroup
	wg.Add(1)
	mSort(arr, &res, &wg)
	wg.Wait()
	return
}

func mSort(arr []int, dest *[]int, wg *sync.WaitGroup) {
	defer wg.Done()
	if len(arr) <= 1 {
		*dest = append(*dest, arr...)
		return
	}

	var left, right []int

	var waitG sync.WaitGroup
	waitG.Add(2)
	go mSort(arr[:len(arr)/2], &left, &waitG)
	go mSort(arr[len(arr)/2:], &right, &waitG)
	waitG.Wait()

	*dest = merge(left, right)
	return
}

func merge(a, b []int) []int {
	res := make([]int, len(a)+len(b))
	ia, ib, i := 0, 0, 0
	for ia < len(a) && ib < len(b) {
		if a[ia] < b[ib] {
			res[i] = a[ia]
			ia++
		} else {
			res[i] = b[ib]
			ib++
		}
		i++
	}

	for ia < len(a) {
		res[i] = a[ia]
		ia++
		i++
	}

	for ib < len(b) {
		res[i] = b[ib]
		ib++
		i++
	}

	return res
}
