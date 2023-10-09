CREATE TABLE "pedido" (
  "pedido_id" SERIAL PRIMARY KEY,
  "data_pedido" timestamptz NOT NULL default (now()),
  "status_pedido" VARCHAR NOT NULL,
  "valor_pedido" FLOAT NOT NULL,
  "quantidade" INTEGER NOT NULL,
  "cliente_id" INTEGER REFERENCES cliente("cliente_id"),
  "vendedor_id" INTEGER REFERENCES vendedor("vendedor_id")
);