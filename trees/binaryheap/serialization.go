package binaryheap

import "github.com/welllog/data-structures/containers"

func assertSerializationImplementation() {
    var _ containers.JSONSerializer = (*Heap)(nil)
    var _ containers.JSONDeserializer = (*Heap)(nil)
}

func (h *Heap) ToJSON() ([]byte, error) {
    return h.list.ToJSON()
}

func (h *Heap) FromJSON(data []byte) error {
    return h.list.FromJSON(data)
}
