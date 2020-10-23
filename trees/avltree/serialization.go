package avltree

import (
    "encoding/json"
    "github.com/welllog/data-structures/containers"
    "github.com/welllog/data-structures/utils"
)

func assertSerializationImplementation() {
    var _ containers.JSONSerializer = (*Tree)(nil)
    var _ containers.JSONDeserializer = (*Tree)(nil)
}

func (tree *Tree) ToJSON() ([]byte, error) {
    elements := make(map[string]interface{})
    it := tree.Iterator()
    for it.Next() {
        elements[utils.ToString(it.Key())] = it.Value()
    }
    return json.Marshal(&elements)
}

func (tree *Tree) FromJSON(data []byte) error {
    elements := make(map[string]interface{})
    err := json.Unmarshal(data, &elements)
    if err == nil {
        tree.Clear()
        for key, value := range elements {
            tree.Put(key, value)
        }
    }
    return err
}
