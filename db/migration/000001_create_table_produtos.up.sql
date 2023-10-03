CREATE TABLE "produtos" (
  "ProdutoID" SERIAL PRIMARY KEY NOT NULL,
  "NomeProduto" VARCHAR NOT NULL,
  "ValorProduto" FLOAT NOT NULL,
  "Estoque" INTEGER NOT NULL
);
