package main

func main() {
	db := newBd("Produtos.csv")
	if err := db.SalvarAllStruct(produtos); err != nil {
		panic(err.Error())
	}
	if err := db.ListAll(); err != nil {
		panic(err.Error())
	}

	db.Close()
}
