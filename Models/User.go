package Models

type User struct {
	ID       int    `gorm:"column:id_user;primaryKey;autoIncrement"  json:"id_user"`
	Username string `gorm:"column:username;type:varchar(255);unique" json:"username"`
	Password string `gorm:"column:password;type:varchar(255)"        json:"password"`
	Role     string `gorm:"column:role;type:enum('admin','kasir')"   json:"role"`
}

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type UserListResponse struct {
	Users []UserResponse `json:"users"`
}