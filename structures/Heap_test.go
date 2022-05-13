package structures

import (
	"LeetCode/structures"
	"container/heap"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_intHeap(t *testing.T)  {
	ast := assert.New(t)

	var ih = new(structures.IntHeap)
	heap.Init(ih)

	begin, end := 0, 10
	for i := begin; i < end; i++ {
		heap.Push(ih, i)
		ast.Equal(0, (*ih)[0], "插入 %d 后的最小值却是 %d，ih=%v", i, (*ih)[0], (*ih))
	}

	for i := begin; i < end; i++ {
		fmt.Println(i, *ih)
		ast.Equal(i, heap.Pop(ih), "Pop 后 ih=%v", (*ih))
	}
}