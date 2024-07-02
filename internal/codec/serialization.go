package codec

// 序列化/反序列化工具接口
type Serializer interface {
	Unmarshal([]byte, interface{}) error
	Marshal(interface{}) ([]byte, error)
}
