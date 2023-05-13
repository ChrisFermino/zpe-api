package services

import (
    "github.com/dlclark/regexp2"
    "zpeTest/src/database"
    "zpeTest/src/models"
)

func CreateUser(user models.User) (any, string) {
    rePassword := regexp2.MustCompile(`^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[$*&@#])[0-9a-zA-Z$*&@#]{8,}$`, 0)
    matchPass, _ := rePassword.MatchString(user.Password)
    if matchPass == false {
        return "", "password strength is not enough"
    }

    uDB, _ := findByEmail(user.Email)
    if uDB.ID != 0 {
        return "", "there is already a user with this email"
    }

    ctxDB := database.DB.Save(&user)
    if ctxDB.Error != nil {
        return "", ctxDB.Error.Error()
    }
    return user, ""
}

func findByEmail(email string) (models.User, error) {
    var u models.User
    if err := database.DB.Where("email LIKE ?", "%"+email+"%").First(&u); err != nil {
        return u, err.Error
    }
    return u, nil
}
