package services

import (
	"github.com/dlclark/regexp2"
	"reflect"
	"zpeTest/src/database"
	"zpeTest/src/models"
	"zpeTest/src/utils"
)

// Service layer with User business rules

func FindUserByFilters(filters map[string]string) ([]models.UserDTo, error) {
	var users []models.UserDTo
	query := database.DB
	for field, value := range filters {
		if value != "" {
			query = query.Where(field+" = ?", value)
		}
	}
	if err := query.Model(&models.User{}).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func findUserById(id string) (models.User, error) {
	var user models.User
	if ctxDB := database.DB.Model(&models.User{}).First(&user, id); ctxDB != nil {
		return user, ctxDB.Error
	}
	return user, nil
}

func findByEmail(email string) (models.User, error) {
	var u models.User
	if err := database.DB.Where("email = ?", "%"+email+"%").First(&u); err != nil {
		return u, err.Error
	}
	return u, nil
}

func CreateUser(user models.User) (any, string) {
	rePassword := regexp2.MustCompile(`^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[$*&@#])[0-9a-zA-Z$*&@#]{8,}$`, 0)
	matchPass, _ := rePassword.MatchString(user.Password)
	if matchPass == false {
		return "", "password strength is not enough"
	}

	userDB, _ := findByEmail(user.Email)
	if userDB.ID != 0 {
		return "", "there is already a user with this email"
	}
	user.Password, _ = utils.HashPassword(user.Password)
	ctxDB := database.DB.Save(&user)
	if ctxDB.Error != nil {
		return "", ctxDB.Error.Error()
	}
	return user, ""
}

func editUser(user interface{}, id string) (string, string) {
	userValue := reflect.ValueOf(user)
	emailValue := userValue.FieldByName("Email")
	userEmail := emailValue.String()

	userDB, _ := findByEmail(userEmail)
	if userDB.ID != 0 {
		return "", "there is already a user with this email"
	}

	userDB, err := findUserById(id)
	if err != nil {
		return "", err.Error()
	}
	ctxDB := database.DB.Model(&userDB).UpdateColumns(user)
	if ctxDB.Error != nil {
		return "", ctxDB.Error.Error()
	}
	return "User updated successfully", ""
}

func EditUserDTO(user models.UserUpdateDTO, id string) (string, string) {
	return editUser(user, id)
}

func EditUserWithRoleDTO(user models.UserUpdateRoleDTO, id string) (string, string) {
	return editUser(user, id)
}

func DeleteUser(id string) (any, string) {
	user, err := findUserById(id)
	if err != nil {
		return nil, err.Error()
	}
	if ctxDB := database.DB.Model(&user).Delete(user); ctxDB.Error != nil {
		return nil, ctxDB.Error.Error()
	}
	return "User deleted successfully", ""
}
