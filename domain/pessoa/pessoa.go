package pessoa

type Pessoa struct {
	Id             string `json:"Id,omitempty"`
	Nome           string `json:"Nome"`
	Cpf            string `json:"Cpf"`
	Datanascimento string `json:"Datanascimento"`
	Email          string `json:"Email"`
	Endereco       string `json:"Endereco"`
	Password       string `json:"Password" validate:"required"`
}

func NewUser(
	Id string,
	Nome string,
	Cpf string,
	Datanascimento string,
	Password string,
) *Pessoa {
	return &Pessoa{
		Id:             Id,
		Cpf:            Cpf,
		Datanascimento: Datanascimento,
		Nome:           Nome,
		Password:       Password,
	}

}
