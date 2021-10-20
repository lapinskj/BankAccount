package entities

type BankAccount struct {
	AccountID 		int32	`gorm:"type:int not null;primaryKey"`
	OwnerName 		string	`gorm:"type:varchar(40) not null"`
	OwnerSurname 	string	`gorm:"type:varchar(40) not null"`
	Balance 		float64	`gorm:"not null"`
}
