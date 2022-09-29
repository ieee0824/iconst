package iconst

func Const[T any](v T) *iconst[T] {
	return &iconst[T]{
		val: v,
	}
}

type iconst[T any] struct {
	val T
}

func (impl *iconst[T]) V() T {
	return impl.val
}
