CREATE TABLE "produtos_pedidos" (
  "ProdutoID" INTEGER NOT NULL,
  "PedidoID" INTEGER NOT NULL,
  PRIMARY KEY ("ProdutoID", "PedidoID"),
  FOREIGN KEY ("ProdutoID") REFERENCES produtos("ProdutoID"),
  FOREIGN KEY ("PedidoID") REFERENCES pedidos("PedidoID")
);