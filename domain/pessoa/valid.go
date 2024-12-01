package pessoa

func (pesssoa Pessoa) Valid() bool {
	if len(pesssoa.Nome) == 0 || pesssoa.Nome == "" {
		return false
	}

	if len(pesssoa.Cpf) == 0 || pesssoa.Cpf == "" {
		return false
	}

	if len(pesssoa.Password) == 0 || pesssoa.Password == "" {
		return false
	}

	return true
}
