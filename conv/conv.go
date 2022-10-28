package conv

type Marshaler[T any] interface {
	Marshal() (T, error)
}

func Marshal[T any](v Marshaler[T]) (T, error) {
	return v.Marshal()
}

type Unmarshaler[T any] interface {
	Unmarshal(T) error
}

func Unmarshal[T any](data T, v Unmarshaler[T]) error {
	return v.Unmarshal(data)
}
