package livroAlugado

type LivroAlugado struct {
	isbn        string `json:"isbn"`
	codigocopia string `json:"codigocopia"`
	pessoa_id   string `json:"pessoa_id"`
}


func NewLivro(
	id string,
	isbn string,
	codigocopia string,
) *LivroAlugado {
	return &LivroAlugado{
		pessoa_id: id,
		isbn: isbn,
		codigocopia: codigocopia,
	}

}
