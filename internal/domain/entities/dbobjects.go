package entities

type User struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	Role     string `db:"role"`
	RegDate  string `db:"registration_date"`
}

type Context struct {
	ID      int    `db:"id"`
	RegDate string `db:"registration_date"`
	Text    string `db:"text"`
}

type Property struct {
	ID   int    `db:"id"`
	Name string `db:"property"`
}

type Model struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type ModelElement struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Unit struct {
	ID      int    `db:"id"`
	ModelID int    `db:"model_id"`
	RegDate string `db:"registration_date"`
	Text    string `db:"text"`
}
