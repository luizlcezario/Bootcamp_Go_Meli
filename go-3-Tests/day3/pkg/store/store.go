package store

const (
	FileType uint = iota
)

type Store interface {
	Read() (interface{}, error)
	Write(data interface{}) error
}

func NewStore[T interface{}](store uint, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore[T]{fileName}
	}
	return nil
}
