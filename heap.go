package cola_prioridad

import (
	"fmt"
)

type heap[T comparable] struct {
	datos    [10]T
	cantidad int
	cmp      func(T, T) int
}

func CrearHeap[T comparable](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.datos = [10]T{}
	heap.cmp = funcion_cmp
	return heap
}

func (h *heap[T]) EstaVacia() bool {
	return h.cantidad == 0
}

func (h *heap[T]) Encolar(elemento T) {
	h.datos[h.cantidad] = elemento
	h.cantidad++
	h.upheap(&(h.datos), h.cantidad-1)
	fmt.Println(h.datos)
}

func (h *heap[T]) upheap(datos *[10]T, pos_hijo int) {
	if pos_hijo <= 0 {
		return
	}
	pos_padre := (pos_hijo - 1) / 2
	if h.cmp(datos[pos_padre], datos[pos_hijo]) > 0 {
		datos[pos_padre], datos[pos_hijo] = datos[pos_hijo], datos[pos_padre]
		h.upheap(datos, pos_padre)
	}
}

func (h *heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic("La cola está vacia")
	}
	return h.datos[0]
}

func (h *heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic("La cola está vacia")
	}
	h.datos[0], h.datos[h.cantidad-1] = h.datos[h.cantidad-1], h.datos[0]
	h.cantidad--
	fmt.Println(h.datos)
	h.downheap(&(h.datos), 0)
	return h.datos[h.cantidad]
}

func (h *heap[T]) downheap(datos *[10]T, pos_padre int) {
	if pos_padre >= h.cantidad-1 {
		return
	}
	pos_hijo_izq := 2*pos_padre + 1
	pos_hijo_der := 2*pos_padre + 2
	if pos_hijo_izq >= h.cantidad || pos_hijo_der >= h.cantidad {
		return
	}
	pos_reemplazo := h.buscarReemplazo(datos, pos_padre, pos_hijo_izq, pos_hijo_der)
	fmt.Println(pos_reemplazo)
	if datos[pos_reemplazo] != datos[pos_padre] {
		datos[pos_padre], datos[pos_reemplazo] = datos[pos_reemplazo], datos[pos_padre]
		fmt.Println(h.datos)
		h.downheap(datos, pos_reemplazo)
	}

}

func (h *heap[T]) buscarReemplazo(datos *[10]T, pos_padre, pos_hijo_izq, pos_hijo_der int) int {
	if h.cmp(datos[pos_hijo_izq], datos[pos_hijo_der]) < 0 {
		return pos_hijo_izq
	}
	return pos_hijo_der
}

func (h *heap[T]) Cantidad() int {
	return h.cantidad
}
