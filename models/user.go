package models

// USER REGISTRATION
type User struct {
	//id, nombre, apellido, edad, genero, email y password
	Id       int32  `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
	Genero   string `json:"genero"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Rol      string `gorm:"NOT NULL; default:'client'" json:"rol"`
	//usr pic
}

// LOGIN DATA
type UserLogin struct {
	//id, nombre, apellido, edad, genero, email y password
	Email    string `json:"email"`
	Password string `json:"password"`
}
