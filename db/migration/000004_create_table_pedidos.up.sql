CREATE TABLE "pedidos" (
  "PedidoID" SERIAL PRIMARY KEY NOT NULL,
  "DataPedido" timestamptz NOT NULL default (now()),
  "StatusPedido" VARCHAR NOT NULL,
  "ValorPedido" FLOAT NOT NULL,
  "Quantidade" INTEGER NOT NULL,
  "ClienteID" INTEGER REFERENCES clientes("ClienteID"),
  "VendedorID" INTEGER REFERENCES vendedores("VendedorID")
);