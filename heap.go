package cola_prioridad

import (
	"fmt"
)

const _CAPACIDAD_INICIAL = 10
const _VECES_A_AUMENTAR = 2
const _VECES_A_REDUCIR = 2
const _VALOR_PARA_REDUCIR = 4

type heap[T comparable] struct {
	datos    []T
	cantidad int
	cmp      func(T, T) int
}

func CrearHeap[T comparable](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.datos = make([]T, _CAPACIDAD_INICIAL)
	heap.cmp = funcion_cmp
	return heap
}

func CrearHeapArr[T comparable](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.datos = arreglo
	heap.cmp = funcion_cmp
	heap.cantidad = len(arreglo)
	for i := len(heap.datos) - 1; i >= 0; i-- {
		downheap(heap.datos, i, heap.cmp)
	}
	return heap
}

func HeapSort[T comparable](elementos []T, funcion_cmp func(T, T) int) {
	for i := len(elementos) - 1; i >= 0; i-- {
		downheap(elementos, i, funcion_cmp)
	}
	fmt.Println(elementos)
	for i := 0; i < len(elementos); i++ {
		elementos[0], elementos[len(elementos)-1-i] = elementos[len(elementos)-1-i], elementos[0]
		downheap(elementos[:len(elementos)-1-i], 0, funcion_cmp)
	}
}

// Métodos de ColaPrioridad

func (h *heap[T]) EstaVacia() bool {
	return h.cantidad == 0
}

func (h *heap[T]) Encolar(elemento T) {
	if h.cantidad == cap(h.datos) {
		h.redimensionar(cap(h.datos) * _VECES_A_AUMENTAR)
	}
	h.datos[h.cantidad] = elemento
	h.cantidad++
	upheap(h.datos, h.cantidad-1, h.cmp)
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
	if h.cantidad <= cap(h.datos)/_VALOR_PARA_REDUCIR && cap(h.datos) > _CAPACIDAD_INICIAL {
		h.redimensionar(cap(h.datos) / _VECES_A_REDUCIR)
	}
	h.datos[0], h.datos[h.cantidad-1] = h.datos[h.cantidad-1], h.datos[0]
	h.cantidad--
	downheap(h.datos, 0, h.cmp)
	return h.datos[h.cantidad]
}

func (h *heap[T]) Cantidad() int {
	return h.cantidad
}

// Métodos / funciones auxiliares

func upheap[T comparable](datos []T, pos_hijo int, func_cmp func(T, T) int) {
	if pos_hijo <= 0 {
		return
	}
	pos_padre := (pos_hijo - 1) / 2
	if func_cmp(datos[pos_padre], datos[pos_hijo]) < 0 {
		datos[pos_padre], datos[pos_hijo] = datos[pos_hijo], datos[pos_padre]
		upheap(datos, pos_padre, func_cmp)
	}
}

func downheap[T comparable](datos []T, pos_padre int, func_cmp func(T, T) int) {
	if pos_padre >= len(datos) {
		return
	}
	pos_hijo_izq := 2*pos_padre + 1
	pos_hijo_der := 2*pos_padre + 2
	if pos_hijo_izq >= len(datos) && pos_hijo_der >= len(datos) {
		return
	}
	var pos_reemplazo int
	if pos_hijo_der >= len(datos) {
		pos_reemplazo = pos_hijo_izq
	} else {
		pos_reemplazo = buscarReemplazo(datos, pos_padre, pos_hijo_izq, pos_hijo_der, func_cmp)
	}
	if datos[pos_reemplazo] != datos[pos_padre] {
		datos[pos_padre], datos[pos_reemplazo] = datos[pos_reemplazo], datos[pos_padre]
		downheap(datos, pos_reemplazo, func_cmp)
	}
}

func buscarReemplazo[T comparable](datos []T, pos_padre, pos_hijo_izq, pos_hijo_der int, func_cmp func(T, T) int) int {
	if func_cmp(datos[pos_hijo_der], datos[pos_padre]) > 0 && func_cmp(datos[pos_hijo_izq], datos[pos_hijo_der]) <= 0 {
		return pos_hijo_der
	} else if func_cmp(datos[pos_hijo_izq], datos[pos_padre]) > 0 && func_cmp(datos[pos_hijo_izq], datos[pos_hijo_der]) > 0 {
		return pos_hijo_izq
	}
	return pos_padre
}

func (h *heap[T]) redimensionar(nuevaCapacidad int) {
	nueva := make([]T, nuevaCapacidad)
	copy(nueva, h.datos)
	h.datos = nueva
}
