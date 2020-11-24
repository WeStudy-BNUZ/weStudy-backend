package db

type College struct{
	ID int `gorm:"primary_key" json:"id"`
	Name string `json:"name" gorm:"not null"`
}

type Major struct{
	ID int `gorm:"primary_key" json:"id"`
	Name string `json:"name" gorm:"not null"`
	CollageID int `json:"collage_id"`
}
