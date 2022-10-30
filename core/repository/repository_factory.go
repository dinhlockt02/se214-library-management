package repository

type RepositoryFactory interface {
	createDocGiaRepository() DocGiaRepository
	createLoaiDocGiaRepository() LoaiDocGiaRepository
	createThamSoRepository() ThamSoRepository
}
