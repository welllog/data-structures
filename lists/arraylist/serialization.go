package arraylist

import (
    "encoding/json"
    "github.com/welllog/data-structures/containers"
)

func assertSerializationImplementation() {
    var _ containers.JSONSerializer = (*List)(nil)
    var _ containers.JSONDeserializer = (*List)(nil)
}

func (l *List) ToJSON() ([]byte, error) {
    return json.Marshal(l.elements[:l.size])
}

func (l *List) FromJSON(data []byte) error {
    err := json.Unmarshal(data, &l.elements)
    if err == nil {
        l.size = len(l.elements)
    }
    return err
}