package external

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"yudegaki.github.com/rewrite-judge/internal/repositories"
)

func Initialize() {
	if exist := DB.Migrator().HasTable(&repositories.User{}); !exist {
		userInitialize()
	}
}

func userInitialize() {
	// Delete table
	DB.Migrator().DropTable(&repositories.User{})
	// Create table
	DB.AutoMigrate(&repositories.User{})
	// Initialize table
	for i := 1; i < 20; i++ {
		password := fmt.Sprintf("password%d", i)
		encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		tmp := repositories.User{
			Model:             gorm.Model{ID: uint(i)},
			Name:              fmt.Sprintf("user%d", i),
			EncryptedPassword: string(encryptedPassword),
		}
		fmt.Println("this is result", tmp)
		DB.Create(&tmp)
	}
}
