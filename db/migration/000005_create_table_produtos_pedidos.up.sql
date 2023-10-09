CREATE TABLE "produto_pedido" (
  "produto_id" INTEGER NOT NULL,
  "pedido_id" INTEGER NOT NULL,
  PRIMARY KEY ("produto_id", "pedido_id"),
  FOREIGN KEY ("produto_id") REFERENCES produto("produto_id"),
  FOREIGN KEY ("pedido_id") REFERENCES pedido("pedido_id")
);