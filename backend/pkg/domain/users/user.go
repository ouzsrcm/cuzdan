package users

import (
	"immortality/database"
)

func SetupDatabase() {

	dst := GetInterfaces()

	MigrateDatabase(dst...)

	SeedingDatabase()
}

func SeedingDatabase() {

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	SeedUser(db)

}

func GetInterfaces() []interface{} {
	models := make([]interface{}, 0)
	return append(models, &User{}, &Credential{}, &CredentialChange{}, &UserToken{})
}

func MigrateDatabase(dst ...interface{}) error {
	if dst != nil {
		dst = append(dst, GetInterfaces()...)
	}
	res := database.Migrate(dst...)
	return res
}
