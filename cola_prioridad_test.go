package cola_prioridad_test

import (
	TDAHeap "cola_prioridad"
	//"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	/**heap := TDAHeap.CrearHeap(func(a, b int) int { return a - b })
	heap.Encolar(2)
	heap.Encolar(9)
	heap.Encolar(5)
	heap.Encolar(12)
	heap.Encolar(11)
	heap.Encolar(10)
	heap.Desencolar()
	heap.Desencolar()
	arr := []int{4, 2, 3, 2, 8, 6, 7, 8, 10}
	heapArr := TDAHeap.CrearHeapArr(arr, func(a, b int) int { return a - b })
	fmt.Println(heapArr.VerMax())
	**/
	arr2 := []int{4, 2, 3, 2, 8, 6, 7, 8, 11, 2222}

	TDAHeap.HeapSort(arr2, func(a, b int) int { return a - b })
}
