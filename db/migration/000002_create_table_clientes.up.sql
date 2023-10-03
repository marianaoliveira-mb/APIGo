CREATE TABLE "clientes" (
  "ClienteID" SERIAL PRIMARY KEY NOT NULL,
  "NomeCliente" VARCHAR NOT NULL,
  "TelefoneCliente" VARCHAR NOT NULL,
  "Saldo" FLOAT NOT NULL
);