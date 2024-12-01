package pessoa

type PessoaRepository interface {
	Get() ([]Pessoa, error)
	GetById(id string) (Pessoa, error)
	GetByEmail(email string) (Pessoa, error)
	Update(data Pessoa) error
	Create(data Pessoa) error
	Delete(id string) error
}
