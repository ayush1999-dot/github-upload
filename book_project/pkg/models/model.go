package model

//db
type Book struct {
	ID         int    `db:"id"         json:"id"`
	Book_name  string `db:"book_name"  json:"book_name"`
	Author_Id  int    `db:"author_id"  json:"author_id"`
	Book_url   string `db:"book_url"   json:"book_url"`
	Is_deleted string `db:"is_deleted" json:"is_deleted"`
}
type Inventory struct {
	Book_ID  int `db:"book_id"  json:"book_id"`
	ID       int `db:"id"  json:"id"`
	Quantity int `db:"quantity_in_stock" json:"quantity_in_stock"`
	Price    int `db:"price" json:"price"`
}

type Author struct {
	Author_Id   int    `db:"author_id"  json:"author_id"`
	Author_name string `db:"author_name"  json:"author_name"`
}
type Display struct {
	ID        int    `db:"id"         json:"id"`
	Book_name string `db:"book_name"  json:"book_name"`
	Book_url  string `db:"book_url"   json:"book_url"`
}
type Add struct {
	Book_name   string `db:"book_name"  json:"book_name"`
	Author_Id   int    `db:"author_id"  json:"author_id"`
	Author_name int    `db:"author_name"  json:"author_name"`
	Book_url    string `db:"book_url"   json:"book_url"`
	Quantity    int    `db:"quantity_in_stock" json:"quantity_in_stock"`
	Price       int    `db:"price" json:"price"`
}
