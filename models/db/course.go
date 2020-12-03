package db

type Course struct{
	ID int `gorm:"primary_key" json:"id"`

	Name string `json:"name" gorm:"not null"`

	Kind int `json:"kind"`

	Class int `json:"class"`
}
