CREATE TABLE "produto" (
  "produto_id" SERIAL PRIMARY KEY,
  "nome_produto" VARCHAR(60) NOT NULL,
  "valor_produto" FLOAT NOT NULL,
  "estoque" INTEGER NOT NULL
);