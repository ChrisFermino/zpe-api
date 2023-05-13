package database

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "zpeTest/src/models"
    "zpeTest/src/utils"
)

var (
    DB  *gorm.DB
    err error
)

// DbConnect Database connection
func DbConnect() {
    dsn := "host=" + utils.EnvVariable("POSTGRES_HOST") +
        " user=" + utils.EnvVariable("POSTGRES_USER") +
        " password=" + utils.EnvVariable("POSTGRES_PASSWORD") +
        " dbname=" + utils.EnvVariable("POSTGRES_NAME") +
        " port=" + utils.EnvVariable("POSTGRES_PORT") +
        " sslmode=disable"

    DB, err = gorm.Open(postgres.New(postgres.Config{
        DSN:                  dsn,
        PreferSimpleProtocol: true,
    }), &gorm.Config{})

    if err != nil {
        log.Panic("Error when trying to connect to the database")
    }

    err := DB.AutoMigrate(&models.User{})
    if err != nil {
        log.Panic("Error when trying to migrate User data")
    }

    if !HasSeedData() {
        SeedDatabase()
    }
}

// HasSeedData checks if the seed data already exists in the database
func HasSeedData() bool {
    var count int64
    DB.Model(&models.User{}).Count(&count)
    return count > 0
}

// SeedDatabase seeds the database with initial data
func SeedDatabase() {
    hashPassword, err := utils.HashPassword(utils.EnvVariable("Password_SEED"))
    user := models.User{
        Name:     utils.EnvVariable("Name_SEED"),
        Email:    utils.EnvVariable("Email_SEED"),
        Password: hashPassword,
        Role:     utils.EnvVariable("Role_SEED"),
    }
    err = DB.Create(&user).Error
    if err != nil {
        log.Panic("Error when seeding User data")
    }
    log.Println("Database seeding completed")
}
