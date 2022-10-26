package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migarate the schema
	database.AutoMigrate(
		&Equipment{}, 
		&User{}, 
		&Employee{}, 
		&BorrowList{},
	)

	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("qwerty"), 14)

	//setup User
	User1 := User{
		Prefix:  	"Ms.",
		FirstName: "Namjai",
		LastName:  "Meedee",
		Email:      "namn@gmail.com",
		Password:   string(password),
		Address:    "12 Moo 6 ...",
		Birthday:   time.Now(),
		Gender:  	"Female",
		Personal:   "1234567890098",
		Mobile:     "0812345678",
		ProvinceID:   23,
	}
	db.Model(&User{}).Create(&User1)

	User2 := User{
		Prefix:  	"Mr.",
		FirstName: "Terdsak",
		LastName:  "Srasai",
		Email:      "tesra@gmail.com",
		Password:   string(password),
		Address:    "122 Mueng ...",
		Birthday:   time.Now(),
		Gender:  	"Male",
		Personal:   "5355201720164",
		Mobile:     "0782256613",
		ProvinceID:   70,
	}
	db.Model(&User{}).Create(&User2)

	User3 := User{
		Prefix:  	"Mr.",
		FirstName: "Arnon",
		LastName:  "Borriboon",
		Email:      "Arnoo@gmail.com",
		Password:   string(password),
		Address:    "11 Moo 5 ...",
		Birthday:   time.Now(),
		Gender:  	"Male",
		Personal:   "1234567890123",
		Mobile:     "0812345678",
		ProvinceID:   23,
	}
	db.Model(&User{}).Create(&User3)

	Employee1 := Employee{
		//EmployeeID:  599323,
		FirstName: "Manee",
		LastName:  "Srisuruk",
		Email:      "namee@gmail.com",
		Password:   string(password),
	}
	db.Model(&Employee{}).Create(&Employee1)

	Employee2 := Employee{
		//EmployeeID:  599323,
		FirstName: "Suree",
		LastName:  "Junsri",
		Email:      "reee@gmail.com",
		Password:   string(password),
	}
	db.Model(&Employee{}).Create(&Employee2)

	Employee3 := Employee{
		//EmployeeID:  599323,
		FirstName: "Rakmak",
		LastName:  "Teesoot",
		Email:      "love@gmail.com",
		Password:   string(password),
	}
	db.Model(&Employee{}).Create(&Employee3)

	Equipment1 := Equipment{
		Name: 			"Projector",
		Amount: 		15,
		Time: 			time.Now(),
		CategoryID: 	1,
		UnitID: 		1,
		Employee: 		Employee1,
	}
	db.Model(&Equipment{}).Create(&Equipment1)

	Equipment2 := Equipment{
		Name: 			"Whiteboard Pen",
		Amount: 		23,
		Time: 			time.Now(),
		CategoryID: 	2,
		UnitID: 		2,
		Employee: 		Employee2,
	}
	db.Model(&Equipment{}).Create(&Equipment2)

	Equipment3 := Equipment{
		Name: 			"Chairs",
		Amount: 		10,
		Time: 			time.Now(),
		CategoryID: 	3,
		UnitID: 		3,
		Employee: 		Employee2,
	}
	db.Model(&Equipment{}).Create(&Equipment3)



	//setup
	// BorrowList1 := BorrowList{
	// 	BorrowTime: time.Now(),
	// 	Amount: 	2,
	// 	Employee: 	Employee1,
	// 	Member: 	User1,
	// 	Equipment:  Equipment1,
	// }
	// db.Model(&BorrowList{}).Create(&BorrowList1)

	// BorrowList2 := BorrowList{
	// 	BorrowTime: time.Now(),
	// 	Amount: 	1,
	// 	Employee: 	Employee2,
	// 	Member: 	User1,
	// 	Equipment:  Equipment2,
	// }
	// db.Model(&BorrowList{}).Create(&BorrowList2)

}
