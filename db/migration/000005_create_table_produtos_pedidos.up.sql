CREATE TABLE "produtos_pedidos" (
  "produto_id" INTEGER NOT NULL,
  "pedido_id" INTEGER NOT NULL,
  PRIMARY KEY ("produto_id", "pedido_id"),
  FOREIGN KEY ("produto_id") REFERENCES produtos("produto_id"),
  FOREIGN KEY ("pedido_id") REFERENCES pedidos("pedido_id")
);