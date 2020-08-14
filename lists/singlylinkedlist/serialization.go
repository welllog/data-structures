package singlylinkedlist

import (
    "encoding/json"
    "github.com/welllog/data-structures/containers"
)

func assertSerializationImplementation() {
    var _ containers.JSONSerializer = (*List)(nil)
    var _ containers.JSONDeserializer = (*List)(nil)
}

func (l *List) ToJSON() ([]byte, error) {
    return json.Marshal(l.Values())
}

func (l *List) FromJSON(data []byte) error {
    elements := []interface{}{}
    err := json.Unmarshal(data, &elements)
    if err == nil {
        l.Clear()
        l.Add(elements...)
    }
    return err
}