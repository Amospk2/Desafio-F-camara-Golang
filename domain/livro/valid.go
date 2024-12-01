package livro

func (livro Livro) Valid() bool {
	if len(livro.autor) == 0 || livro.autor == "" {
		return false
	}

	if len(livro.titulo) == 0 || livro.titulo == "" {
		return false
	}

	if len(livro.isbn) == 0 || livro.isbn == "" {
		return false
	}

	if len(livro.codigocopia) == 0 || livro.codigocopia == "" {
		return false
	}


	return true
}
