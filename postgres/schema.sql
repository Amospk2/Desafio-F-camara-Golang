
CREATE TABLE IF NOT EXISTS pessoa (
    id uuid UNIQUE NOT NULL,
    nome varchar(255) NOT NULL,
    email varchar(255) UNIQUE NOT NULL,
    password varchar(255) NOT NULL,
    datanascimento varchar(255) NOT NULL,
    endereco varchar(255) NOT NULL,
    cpf varchar(255) not NULL
);

