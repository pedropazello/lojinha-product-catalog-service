package serializers

type ISerializer interface {
	Serialize(interface{}) ([]byte, error)
}
