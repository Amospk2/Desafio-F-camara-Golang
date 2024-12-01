package livroAlugado

func (livro LivroAlugado) Valid() bool {
	if len(livro.codigocopia) == 0 || livro.codigocopia == "" {
		return false
	}

	if len(livro.isbn) == 0 || livro.isbn == "" {
		return false
	}

	if len(livro.pessoa_id) == 0 || livro.pessoa_id == "" {
		return false
	}

	return true
}
