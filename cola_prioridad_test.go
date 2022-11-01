package cola_prioridad_test

import (
	TDAHeap "cola_prioridad"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeapVacio(t *testing.T) {
	heap := TDAHeap.CrearHeap(func(a, b int) int { return a - b })

	require.True(t, heap.EstaVacia())
	require.Panics(t, func() { heap.VerMax() })
	require.Panics(t, func() { heap.Desencolar() })
	require.Zero(t, heap.Cantidad())
	require.NotNil(t, heap)
}

func TestFuncionalidadHeap(t *testing.T) {
	heap := TDAHeap.CrearHeap(func(a, b int) int { return a - b })

	heap.Encolar(10)
	heap.Encolar(5)
	heap.Encolar(20)

	require.False(t, heap.EstaVacia())
	require.Equal(t, 3, heap.Cantidad())
	require.Equal(t, 20, heap.VerMax())
	require.Equal(t, 20, heap.Desencolar())

	require.Equal(t, 10, heap.Desencolar())

	require.Equal(t, 5, heap.VerMax())
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 5, heap.Desencolar())
	require.True(t, heap.EstaVacia())
	require.Panics(t, func() { heap.VerMax() })
	require.Panics(t, func() { heap.Desencolar() })

	require.Zero(t, heap.Cantidad())
	heap.Encolar(1)
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 1, heap.VerMax())
	require.Equal(t, 1, heap.Desencolar())
}

func TestHeap(t *testing.T) {
	heap := TDAHeap.CrearHeap(func(a, b int) int { return a - b })
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
	arr2 := []int{4, 2, 3, 2, 8, 6, 7, 8, 11, 2222}

	TDAHeap.HeapSort(arr2, func(a, b int) int { return a - b })
}
