package store

const (
	FileType uint = iota
)

type Store interface {
	Read(data interface{}) (interface{}, error)
	Write(data interface{}) (interface{}, error)
}

func NewStore(store uint, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{fileName}
	}
	return nil
}
