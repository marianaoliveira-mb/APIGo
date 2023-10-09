CREATE TABLE "cliente" (
  "cliente_id" SERIAL PRIMARY KEY,
  "nome_cliente" VARCHAR(60) NOT NULL,
  "telefone_cliente" VARCHAR NOT NULL,
  "saldo" FLOAT NOT NULL
);