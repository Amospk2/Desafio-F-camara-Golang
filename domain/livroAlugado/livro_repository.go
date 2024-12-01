package livroAlugado

type LivroRepository interface {
	Get() ([]LivroAlugado, error)
	GetById(id string) (LivroAlugado, error)
	Update(data LivroAlugado) error
	Create(data LivroAlugado) error
	Delete(id string) error
}
