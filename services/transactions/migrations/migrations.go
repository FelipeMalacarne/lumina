package migrations

import (
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Id                uuid.UUID       `gorm:"type:uuid;primaryKey"`
	UserId            uuid.UUID       `gorm:"type:uuid;not null"`
	TransactionTypeId uuid.UUID       `gorm:"type:uuid;not null"`
	TransactionType   TransactionType `gorm:"type:varchar(255)"`
	LedgerId          uuid.UUID       `gorm:"type:uuid;not null"`
	Ledger            Ledger          `gorm:"foreignKey:LedgerId"`
	Amount            float64         `gorm:"type:float"`
	FitId             uuid.UUID       `gorm:"type:uuid;not null"`
	Memo              string          `gorm:"type:varchar(255)"`
	Categories        []Category      `gorm:"many2many:transaction_categories;"`
	CreatedAt         time.Time       `gorm:"type:timestamp"`
	UpdatedAt         time.Time       `gorm:"type:timestamp"`
}

type Ledger struct {
	gorm.Model
	Id           uuid.UUID     `gorm:"type:uuid;primaryKey"`
	UserId       uuid.UUID     `gorm:"type:uuid;not null"`
	AccountId    uuid.UUID     `gorm:"type:uuid;not null"`
	Account      Account       `gorm:"foreignKey:AccountId"`
	Transactions []Transaction `gorm:"foreignKey:LedgerId"`
	Name         string        `gorm:"type:varchar(255)"`
	Balance      float64       `gorm:"type:float"`
}

type Account struct {
	gorm.Model
	Id        uuid.UUID `gorm:"type:uuid;primaryKey"`
	BankId    uuid.UUID `gorm:"type:uuid;not null"`
	Name      string    `gorm:"type:varchar(255)"`
	Type      string    `gorm:"type:varchar(255)"`
	Ledgers   []Ledger  `gorm:"foreignKey:AccountId"`
	CreatedAt time.Time `gorm:"type:timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp"`
}

type TransactionType struct {
	gorm.Model
	Id                    uuid.UUID              `gorm:"type:uuid;primaryKey"`
	Name                  string                 `gorm:"type:varchar(255)"`
	Transactions          []Transaction          `gorm:"foreignKey:TransactionTypeId"`
	RecurringTransactions []RecurringTransaction `gorm:"foreignKey:TransactionTypeId"`
}

type RecurringTransaction struct {
	gorm.Model
	Id                uuid.UUID       `gorm:"type:uuid;primaryKey"`
	UserId            uuid.UUID       `gorm:"type:uuid;not null"`
	TransactionTypeId uuid.UUID       `gorm:"type:uuid;not null"`
	TransactionType   TransactionType `gorm:"foreignKey:TransactionTypeId"`
	Amount            float64         `gorm:"type:float"`
	Source            string          `gorm:"type:varchar(255)"`
	Frequency         uint8           `gorm:"type:smallint"`
	NextOccurrence    time.Time       `gorm:"type:timestamp"`
	CreatedAt         time.Time       `gorm:"type:timestamp"`
	UpdatedAt         time.Time       `gorm:"type:timestamp"`
}

type Category struct {
	gorm.Model
	Id           uuid.UUID     `gorm:"type:uuid;primaryKey"`
	UserId       uuid.UUID     `gorm:"type:uuid;not null"`
	Budgets      []Budget      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Transactions []Transaction `gorm:"many2many:transaction_categories;"`
	Name         string        `gorm:"type:varchar(255)"`
	CreatedAt    time.Time     `gorm:"type:timestamp"`
	UpdatedAt    time.Time     `gorm:"type:timestamp"`
}

type TransactionCategory struct {
	gorm.Model
	TransactionId uuid.UUID `gorm:"type:uuid;not null"`
	CategoryId    uuid.UUID `gorm:"type:uuid;not null"`
}

type Budget struct {
	gorm.Model
	Id         uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserId     uuid.UUID `gorm:"type:uuid;not null"`
	CategoryId uuid.UUID `gorm:"type:uuid;not null"`
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&Transaction{}, &Ledger{}, &Account{}, &TransactionType{}, &RecurringTransaction{}, &Category{}, &Budget{})
	if err != nil {
		log.Println("Failed to migrate database", err)
	}
}

func DropTables(db *gorm.DB) {
	err := db.Migrator().DropTable(&Transaction{}, &Ledger{}, &Account{}, &TransactionType{}, &RecurringTransaction{}, &Category{}, &Budget{})
	if err != nil {
		log.Println("Failed to drop tables", err)
	}
}
