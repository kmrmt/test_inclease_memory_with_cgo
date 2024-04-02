package main

// #cgo LDFLAGS: -L. -ltest -lstdc++ -lngt
// #include "c.h"
// #include <NGT/Capi.h>
import "C"

func main() {
	const (
		dim  int = 128
		size int = 1000000
	)

	C.header()

	vectors := make([]float32, dim*size)
	for i := 0; i < size; i++ {
		vectors[i*dim] = float32(i)
	}

	e := C.ngt_create_error_object()
	p := C.ngt_create_property(e)
	C.ngt_set_property_edge_size_for_creation(p, 40, e)
	C.ngt_set_property_edge_size_for_search(p, 40, e)
	C.ngt_set_property_dimension(p, C.int(dim), e)
	C.ngt_set_property_object_type_float(p, e)
	C.ngt_set_property_distance_type_inner_product(p, e)

	idx := C.ngt_create_graph_and_tree_in_memory(p, e)

	ids := make([]C.ObjectID, size)
	C.stat(C.CString("init"))
	for i := 0; i < size; i++ {
		ids[i] = C.ngt_insert_index_as_float(idx, (*C.float)(&vectors[i*dim]), C.uint32_t(dim), e)
	}
	C.stat(C.CString("insert"))

	C.ngt_create_index(idx, 16, e)
	C.stat(C.CString("create_index"))

	C.ngt_close_index(idx)
	C.ngt_destroy_property(p)
	C.ngt_destroy_error_object(e)

	C.stat(C.CString("finish"))
}
