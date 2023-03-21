package dateset

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
}

type UpdateUser struct {
	ID           int    `json:"id"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Phone        string `json:"phone"`
	Profile_user string `json:"profile_user"`
}

type LoginUser struct {
	Status string `json:"status"`
	ID     int    `json:"id"`
	TypeNumber_User int `json:"typenumber_user"`
}

type GetUser struct {
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Profile_user string `json:"profile_user"`
}

type AddWorkFreelance struct {
	WorkPostID             int    `json:"workpostid"`
	TypeWorkNumber         int    `json:"typeWorknumber"`
	DetailWork             string `json:"detailwork"`
	UserID                 int    `json:"userid"`
	PricePostWork          string `json:"pricepostwork"`
	NameWork               string `json:"namework"`
	ImageWorkPostFreelance string `json:"imageworkpostfreelance"`
}

type ID struct {
	ID int `json:"id"`
}

type UpdateFreelance struct {
	ID           int    `json:"id"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Profile_User string `json:"profileuser"`
	Line         string `json:"line"`
	Facebook     string `json:"facebook"`
	Instagram    string `json:"instagram"`
}

type GetFreelance struct {
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Profile_User string `json:"profileuser"`
	Line         string `json:"line"`
	Facebook     string `json:"facebook"`
	Instagram    string `json:"instagram"`
}
