package migrations

import (
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id              uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name            string    `gorm:"type:varchar(255)"`
	Email           string    `gorm:"type:varchar(255);uniqueIndex"`
	EmailVerifiedAt *string   `gorm:"type:timestamp"`
	Password        string    `gorm:"type:varchar(255)"`
	RememberToken   string    `gorm:"type:varchar(100)"`
	Members         []Member  `gorm:"foreignKey:UserId"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}

type Member struct {
	gorm.Model
	Id        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserId    uuid.UUID `gorm:"type:uuid;not null"`
	User      User      `gorm:"foreignKey:UserId"`
	ProjectId uuid.UUID `gorm:"type:uuid;not null"`
	Project   Project   `gorm:"foreignKey:ProjectId"`
	RoleId    uuid.UUID `gorm:"type:uuid;not null"`
	Role      Role      `gorm:"foreignKey:RoleId"`
}

type Project struct {
	gorm.Model
	Id      uuid.UUID `gorm:"type:uuid;primaryKey"`
	Members []Member  `gorm:"foreignKey:ProjectId"`
}

type Role struct {
	gorm.Model
	Id         uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Members    []Member   `gorm:"foreignKey:RoleId"`
	Permission Permission `gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Permission struct {
	gorm.Model
	Id              uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name            string    `gorm:"type:varchar(255)"`
	RoleId          uuid.UUID `gorm:"type:uuid;not null"`
	CanAddMember    bool      `gorm:"type:boolean"`
	CanRemoveMember bool      `gorm:"type:boolean"`
}

type PasswordResetToken struct {
	gorm.Model
	Id        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex"`
	Token     string    `gorm:"type:varchar(255);uniqueIndex"`
	CreatedAt time.Time
}

type PersonalAccessToken struct {
	gorm.Model
	Id        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Token     string    `gorm:"type:varchar(255);uniqueIndex"`
	UserId    uuid.UUID `gorm:"type:uuid;not null"`
	User      User      `gorm:"foreignKey:UserId"`
	Name      string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&User{}, &Member{}, &Project{}, &Role{}, &Permission{}, &PasswordResetToken{}, &PersonalAccessToken{})
	if err != nil {
		log.Println("Failed to migrate database", err)
	}
}

func DropTables(db *gorm.DB) {
	err := db.Migrator().DropTable(&User{}, &Member{}, &Project{}, &Role{}, &Permission{}, &PasswordResetToken{}, &PersonalAccessToken{})
	if err != nil {
		log.Println("Failed to drop tables", err)
	}
}
