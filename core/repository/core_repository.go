package repository

type CoreRepository interface {
	startTransaction()
	commitTransaction()
	rollbackTransaction()
}
