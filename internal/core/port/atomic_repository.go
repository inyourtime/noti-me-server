package port

type Repository interface {
	UserRepository() UserRepository
}

type AtomicRepoositoryCallback func(params Repository) error

type AtomicRepository interface {
	Transaction(callback AtomicRepoositoryCallback) error
}
