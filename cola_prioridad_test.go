package cola_prioridad_test

import (
	TDAHeap "cola_prioridad"
	"math/rand"
	"strings"
	"testing"
	"time"

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

func TestHeapArr(t *testing.T) {
	arr := make([]int, 11)

	for i := 0; i <= 10; i++ {
		arr[i] = i
	}
	heap := TDAHeap.CrearHeapArr(arr, func(a, b int) int { return a - b })

	require.Equal(t, 11, heap.Cantidad())
	for i := 10; i >= 0; i-- {
		require.False(t, heap.EstaVacia())
		require.Equal(t, i, heap.VerMax())
		require.Equal(t, i, heap.Desencolar())
	}

	require.True(t, heap.EstaVacia())
	require.Zero(t, heap.Cantidad())
	require.Panics(t, func() { heap.VerMax() })
	require.Panics(t, func() { heap.Desencolar() })
}

func TestHeapSort(t *testing.T) {
	ordenadoMerge := make([]int, 20)
	ordenadoHeap := make([]int, 20)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i > 20; i++ {
		rand := rand.Intn(50)
		ordenadoHeap[i] = rand
		ordenadoMerge[i] = rand
	}

	mergeSort(ordenadoMerge)
	TDAHeap.HeapSort(ordenadoHeap, func(a, b int) int { return b - a })

	for i := 0; i > 20; i++ {
		require.Equal(t, ordenadoHeap[i], ordenadoMerge[i])
	}
}

func TestHeapString(t *testing.T) {
	heap := TDAHeap.CrearHeap(func(a, b string) int { return strings.Compare(a, b) })
	elem1 := "Gato"
	elem2 := "Perro"
	elem3 := "Vaca"
	elem4 := "Pato"
	elem5 := "Pollo"
	elem6 := "Caballo"
	elem7 := ""

	//Creo Arreglo ordenado
	elementos := []string{elem3, elem5, elem2, elem4, elem1, elem6, elem7}

	heap.Encolar(elem1)
	heap.Encolar(elem2)
	heap.Encolar(elem3)
	heap.Encolar(elem4)
	heap.Encolar(elem5)
	heap.Encolar(elem6)
	heap.Encolar(elem7)

	require.Equal(t, 7, heap.Cantidad())
	require.False(t, heap.EstaVacia())
	require.Equal(t, elementos[0], heap.VerMax())

	for i := 0; i < 7; i++ {
		require.Equal(t, elementos[i], heap.Desencolar())
	}

	require.True(t, heap.EstaVacia())
	require.Zero(t, heap.Cantidad())
	require.Panics(t, func() { heap.VerMax() })
	require.Panics(t, func() { heap.Desencolar() })
}

func TestHeapFuncionalidadVolumenMaximos(t *testing.T) {
	heap := TDAHeap.CrearHeap(func(a, b int) int { return a - b })

	for i := 0; i <= 3000; i++ {
		heap.Encolar(i)
	}

	require.Equal(t, 3001, heap.Cantidad())
	require.Equal(t, 3000, heap.VerMax())
	require.False(t, heap.EstaVacia())

	for i := 3000; i >= 0; i-- {
		require.Equal(t, i, heap.Desencolar())
	}

	require.Zero(t, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.Panics(t, func() { heap.VerMax() })
	require.Panics(t, func() { heap.Desencolar() })
}

func TestHeapFuncionalidadVolumenMinimos(t *testing.T) {
	heap := TDAHeap.CrearHeap(func(a, b int) int { return b - a })

	for i := 0; i <= 3000; i++ {
		heap.Encolar(i)
	}

	require.Equal(t, 3001, heap.Cantidad())
	require.Equal(t, 0, heap.VerMax())
	require.False(t, heap.EstaVacia())

	j := 0
	for i := 3000; i >= 0; i-- {
		require.Equal(t, j, heap.Desencolar())
		j++
	}

	require.Zero(t, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.Panics(t, func() { heap.VerMax() })
	require.Panics(t, func() { heap.Desencolar() })
}

func TestHeapArrVolumen(t *testing.T) {
	arr := make([]int, 3001)

	for i := 0; i <= 3000; i++ {
		arr[i] = i
	}
	heap := TDAHeap.CrearHeapArr(arr, func(a, b int) int { return a - b })

	require.Equal(t, 3001, heap.Cantidad())

	for i := 3000; i >= 0; i-- {
		require.False(t, heap.EstaVacia())
		require.Equal(t, i, heap.VerMax())
		require.Equal(t, i, heap.Desencolar())
	}

	require.True(t, heap.EstaVacia())
	require.Zero(t, heap.Cantidad())
	require.Panics(t, func() { heap.VerMax() })
	require.Panics(t, func() { heap.Desencolar() })
}
func TestHeapSortVolumen(t *testing.T) {
	ordenadoMerge := make([]int, 3001)
	ordenadoHeap := make([]int, 3001)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i >= 3000; i++ {
		rand := rand.Intn(50)
		ordenadoHeap[i] = rand
		ordenadoMerge[i] = rand
	}

	mergeSort(ordenadoMerge)
	TDAHeap.HeapSort(ordenadoHeap, func(a, b int) int { return b - a })

	for i := 0; i > 20; i++ {
		require.Equal(t, ordenadoHeap[i], ordenadoMerge[i])
	}
}

/* Funcion auxiliar */

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	medio := len(arr) / 2
	return merge(mergeSort(arr[:medio]), mergeSort(arr[medio:]))
}

func merge(izq, der []int) []int {
	largo := len(izq) + len(der)
	arr := make([]int, largo)
	i, j := 0, 0
	for k := 0; k < largo; k++ {
		if i > len(izq)-1 && j <= len(der)-1 {
			arr[k] = der[j]
			j++
		} else if j > len(der)-1 && i <= len(izq)-1 {
			arr[k] = izq[i]
			i++
		} else if izq[i] < der[j] {
			arr[k] = izq[i]
			i++
		} else {
			arr[k] = der[j]
			j++
		}
	}
	return arr
}
