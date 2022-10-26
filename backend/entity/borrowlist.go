package entity

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model

	FirstName	string
	LastName	string
	Email		string	`gorm:"uniqueIndex"`
	Password	string
	
	Equipment 	[]Equipment		`gorm:"foreignKey:EmployeeID"`
	BorrowList	[]BorrowList	`gorm:"foreignKey:EmployeeID"`
}

type User struct {
	gorm.Model

	//PrefixID	*uint
	Prefix		string	//Prefix `gorm:"references:ID"`

	FirstName	string
	LastName	string
	Email		string	`gorm:"uniqueIndex"`
	Password	string
	Address		string
	Birthday	time.Time

	//GenderID	*uint
	Gender		string	//Gender `gorm:"references:ID"`

	Personal	string	`gorm:"uniqueIndex"`
	Mobile		string	`gorm:"uniqueIndex"`
	
	ProvinceID	int
	//Province	Province	`gorm:"references:ID"`

	BorrowList []BorrowList	`gorm:"foreignKey:MemberID"`
}

type Equipment struct {
	gorm.Model

	Name			string
	Amount			int
	Time	 		time.Time

	CategoryID		int
	// Category		Catagory	`gorm:"references:ID"`

	UnitID			int
	// Unit			Unit		`gorm:"references:ID"`

	EmployeeID		*uint
	Employee		Employee	`gorm:"references:ID"`

	BorrowList []BorrowList		`gorm:"foreignKey:EquipmentID"`
}

type BorrowList	struct {
	gorm.Model

	BorrowTime		time.Time
	Amount			int

	EquipmentID 	*uint		
	Equipment		Equipment	`gorm:"references:ID"`

	MemberID		*uint
	Member			User		`gorm:"references:ID"`

	EmployeeID		*uint
	Employee		Employee	`gorm:"references:ID"`
}