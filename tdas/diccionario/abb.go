package diccionario

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
}

func (a *abb[K, V]) Guardar(clave K, dato V) {
}

func (a *abb[K, V]) Pertenece(clave K) bool {
}

func (a *abb[K, V]) Obtener(clave K) V {
}

func (a *abb[K, V]) Borrar(clave K) V {
}

func (a *abb[K, V]) Cantidad() int {
}

func (a *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
}

func (a *abb[K, V]) Iterador() IterDiccionario[K, V] {
}

func (a *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
}

func (a *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
}
