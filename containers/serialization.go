package containers

// Serialization提供了序列化器（marshalers）和解序列化器（unmarshalers）。

// JSONSerializer 提供JSON序列化
type JSONSerializer interface {
    // ToJSON输出容器元素的JSON表示。
    ToJSON() ([]byte, error)
}

// JSONDeserializer 提供JSON反序列化
type JSONDeserializer interface {
    // FromJSON从输入的JSON表示中填充容器的元素。
    FromJSON([]byte) error
}