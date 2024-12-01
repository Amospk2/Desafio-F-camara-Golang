package livro

type Livro struct {
	id          string `json:"id,omitempty"`
	titulo      string `json:"titulo"`
	autor       string `json:"autor"`
	isbn        string `json:"isbn"`
	codigocopia string `json:"codigocopia"`
}

func NewLivro(
	id string,
	titulo string,
	autor string,
	isbn string,
	codigocopia string,
) *Livro {
	return &Livro{
		titulo: titulo,
		autor: autor,
		isbn: isbn,
		codigocopia: codigocopia,
	}

}
