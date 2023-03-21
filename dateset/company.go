package dateset

type Company struct {
	ID           int    `json:"id"`
	CompanyName  string `json:"companyname"`
	CompanyEmail string `json:"companyemail"`
	CompanyPhone string `json:"companyphone"`
	Address      string `json:"address"`
	Subdistrict  string `json:"subdistrict"`
	District     string `json:"district"`
	Province     string `json:"province"`
	Postcode     string `json:"postcode"`
	Password     string `json:"password"`
}

type AddWorkCompany struct {
	WorkPostID           int    `json:"workpostid"`
	CompanyID            int    `json:"companyid"`
	TypeWorkNumber       int    `json:"typeworknumber"`
	NameWork             string `json:"namework"`
	DetailWork           string `json:"detailwork"`
	Position             string `json:"position"`
	NumPerson            int    `json:"numperson"`
	PriceWorkMin         string `json:"priceworkmin"`
	PriceWorkMax         string `json:"priceworkmax"`
	Education            string `json:"education"`
	ImageWorkPostCompany string `json:"imageworkpostcompany"`
}

type GetCompany struct {
	CompanyName    string `json:"companyname"`
	CompanyEmail   string `json:"companyemail"`
	CompanyPhone   string `json:"companyphone"`
	Address        string `json:"address"`
	Subdistrict    string `json:"subdistrict"`
	District       string `json:"district"`
	Province       string `json:"province"`
	Postcode       string `json:"postcode"`
	ProfileCompany string `json:"profilecompany"`
}
type UpdateCompany struct {
	ID             int    `json:"id"`
	CompanyName    string `json:"companyname"`
	CompanyEmail   string `json:"companyemail"`
	CompanyPhone   string `json:"companyphone"`
	Address        string `json:"address"`
	Subdistrict    string `json:"subdistrict"`
	District       string `json:"district"`
	Province       string `json:"province"`
	Postcode       string `json:"postcode"`
	ProfileCompany string `json:"profilecompany"`
}