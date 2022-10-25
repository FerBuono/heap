package cola_prioridad_test

import (
	TDAHeap "cola_prioridad"
	"testing"
)

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
}
