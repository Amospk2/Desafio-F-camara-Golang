package database

import (
	"context"
	"desafiot/domain/pessoa"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PessoaRepositoryImp struct {
	pool *pgxpool.Pool
}

func (db PessoaRepositoryImp) Get() ([]pessoa.Pessoa, error) {

	PESSOA := make([]pessoa.Pessoa, 0)

	rows, err := db.pool.Query(context.Background(),
		"select id, nome, email, password, cpf, datanascimento, endereco from public.PESSOA",
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user pessoa.Pessoa

		err = rows.Scan(
			&user.Id,
			&user.Nome,
			&user.Email,
			&user.Password,
			&user.Cpf,
			&user.Datanascimento,
			&user.Endereco,
		)

		if err != nil {
			log.Fatal(err)
		}

		PESSOA = append(PESSOA, user)
	}

	return PESSOA, nil
}

func (db PessoaRepositoryImp) GetById(id string) (pessoa.Pessoa, error) {

	var userFinded pessoa.Pessoa

	err := db.pool.QueryRow(
		context.Background(),
		"select id, nome, email, password, cpf, datanascimento, endereco from public.PESSOA where id=$1", id,
	).Scan(
		&userFinded.Id,
		&userFinded.Nome,
		&userFinded.Email,
		&userFinded.Password,
		&userFinded.Cpf,
		&userFinded.Datanascimento,
		&userFinded.Endereco,
	)

	if err != nil {
		return pessoa.Pessoa{}, err
	}

	return userFinded, nil
}

func (db PessoaRepositoryImp) GetByEmail(email string) (pessoa.Pessoa, error) {

	var userFinded pessoa.Pessoa

	err := db.pool.QueryRow(
		context.Background(),
		"select id, nome, email, password, cpf, datanascimento, endereco from public.PESSOA where email=$1", email,
	).Scan(
		&userFinded.Id,
		&userFinded.Nome,
		&userFinded.Email,
		&userFinded.Password,
		&userFinded.Cpf,
		&userFinded.Datanascimento,
		&userFinded.Endereco,
	)

	if err != nil {
		return pessoa.Pessoa{}, err
	}

	return userFinded, nil
}

func (db PessoaRepositoryImp) Update(data pessoa.Pessoa) error {
	_, err := db.pool.Exec(
		context.Background(),
		"UPDATE PESSOA SET nome = $1, email = $2, password = $3, "+
			"datanascimento = $4, endereco = $5, cpf = $6, WHERE id = $7",
		data.Nome, data.Email, data.Password, data.Datanascimento, data.Endereco, data.Cpf, data.Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db PessoaRepositoryImp) Create(data pessoa.Pessoa) error {
	_, err := db.pool.Exec(
		context.Background(), "INSERT INTO PESSOA VALUES($1,$2,$3,$4,$5,$6,$7)",
		data.Id, data.Nome, data.Email, data.Password, data.Datanascimento, data.Endereco, data.Cpf,
	)

	if err != nil {
		return err
	}

	return nil
}

func (db PessoaRepositoryImp) Delete(id string) error {

	_, err := db.pool.Exec(context.Background(), "DELETE FROM PESSOA WHERE id=$1", id)

	if err != nil {
		return err
	}

	return nil
}

func NewPessoaRepositoryImp(pool *pgxpool.Pool) pessoa.PessoaRepository {
	return PessoaRepositoryImp{pool: pool}
}
