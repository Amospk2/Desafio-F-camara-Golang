package livro

type LivroRepository interface {
	Get() ([]Livro, error)
	GetById(id string) (Livro, error)
	Update(data Livro) error
	Create(data Livro) error
	Delete(id string) error
}
