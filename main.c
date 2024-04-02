#include "c.h"
#include <NGT/Capi.h>
#include <stdlib.h>
#include <stdio.h>

int main() {
    const int dim = 128;
    const int size = 1000000;

    header();
    stat("start");

    float* vectors = (float*)malloc(sizeof(float) * dim * size);
    for (int i = 0; i < size; i++) {
        vectors[i * dim] = i;
    }
    ObjectID *ids = (ObjectID*)malloc(sizeof(ObjectID) * size);

    NGTError e = ngt_create_error_object();
    NGTProperty p = ngt_create_property(e);
    ngt_set_property_edge_size_for_creation(p, 40, e);
    ngt_set_property_edge_size_for_search(p, 40, e);
    ngt_set_property_dimension(p, dim, e);
    ngt_set_property_object_type_float(p, e);
    ngt_set_property_distance_type_inner_product(p, e);

    NGTIndex idx = ngt_create_graph_and_tree_in_memory(p, e);

    stat("init");
    for (int i = 0; i < size; i++) {
        ids[i] = ngt_insert_index_as_float(idx, &vectors[i * dim], dim, e);
    }
    stat("insert");

    ngt_create_index(idx, 16, e);
    stat("create_index");

    ngt_close_index(idx);
    ngt_destroy_property(p);
    ngt_destroy_error_object(e);

    free(ids);
    free(vectors);

    stat("finish");
}