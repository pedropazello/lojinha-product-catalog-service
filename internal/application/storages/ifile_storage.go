package storages

type IFileStorage interface {
	GetFile(fileName string) ([]byte, error)
}
