package hashset

import (
    "encoding/json"
    "github.com/welllog/data-structures/containers"
)

func assertSerializationImplementation() {
    var _ containers.JSONSerializer = (*Set)(nil)
    var _ containers.JSONDeserializer = (*Set)(nil)
}

func (set *Set) ToJSON() ([]byte, error) {
    return json.Marshal(set.Values())
}

func (set *Set) FromJSON(data []byte) error {
    elements := []interface{}{}
    err := json.Unmarshal(data, &elements)
    if err == nil {
        set.Clear()
        set.Add(elements...)
    }
    return err
}
