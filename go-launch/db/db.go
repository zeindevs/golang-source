package db

type db struct{}

type queryBuilder[T any] struct{}

func Raw[T any](raw string) *queryBuilder[T] {
	b := &queryBuilder[T]{}
	return b
}

func Where[T any](args ...any) *queryBuilder[T] {
	b := &queryBuilder[T]{}
	return b
}

func (q *queryBuilder[T]) Limit(n int) *queryBuilder[T] {
	return q
}

func (q *queryBuilder[T]) Asc(field string) *queryBuilder[T] {
	return q
}

func (q *queryBuilder[T]) Desc(field string) *queryBuilder[T] {
	return q
}

func (q *queryBuilder[T]) Run() (T, error) {
	var t T
	return t, nil
}
