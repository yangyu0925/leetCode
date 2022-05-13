package main

func merge(A []int, m int, B []int, n int)  {
	AIndex := m - 1
	BIndex := n - 1
	idx := len(A) - 1
	for AIndex >= 0 && BIndex >= 0  {
		if A[AIndex] > B[BIndex] {
			A[idx] = A[AIndex]
			AIndex--
		} else {
			A[idx] = B[BIndex]
			BIndex--
		}
		idx--
	}
	for BIndex >= 0 {
		A[idx]= B[BIndex]
		BIndex--
		idx--
	}
}