package mockdb

import (
	"fmt"
	"github.com/rpupo63/ProNexus/backend/errs"
	"github.com/rpupo63/ProNexus/backend/models"
	"github.com/google/uuid"
    "reflect"
)

type UserRepo struct {
	users *[]models.User
}

func NewUserRepo(users *[]models.User) *UserRepo {
	return &UserRepo{users}
}

func (r *UserRepo) SelectAll() ([]models.User, error) {
	return *r.users, nil
}

func (r *UserRepo) FindByEmail(email string) (models.User, error) {
    for _, user := range *r.users {
        // Correctly compare the passed email with the Email field of the user
        if user.Email == email {
            return user, nil
        }
    }
    // Assuming errs.NewNotFound is a correct call to a custom error handling function
    return models.User{}, errs.NewNotFound("user")
}


func (r *UserRepo) FindByID(id string) (models.User, error) {
    for _, user := range *r.users {
        // Correctly compare the passed email with the Email field of the user
        if user.ID == id {
            return user, nil
        }
    }
    // Assuming errs.NewNotFound is a correct call to a custom error handling function
    return models.User{}, errs.NewNotFound("user")
}



func (r *UserRepo) Insert(desiredUser models.User) error {
	if desiredUser.ID == "" {
		desiredUser.ID = uuid.New().String()
	}
	*r.users = append(*r.users, desiredUser)
	return nil
}

func (r *UserRepo) Update(userFields models.User) error {
	if userFields.ID == "" {
		return fmt.Errorf("error: missing ID field in argument")
	}

	for i, user := range *r.users {
		if user.ID == userFields.ID {
			userValue := reflect.ValueOf(&user).Elem()
			fieldsValue := reflect.ValueOf(userFields)

			for j := 0; j < fieldsValue.NumField(); j++ {
				field := fieldsValue.Field(j)
				if !field.IsZero() {
					userField := userValue.Field(j)
					userField.Set(field)
				}
			}

			(*r.users)[i] = user
			break
		}
	}

	return nil
}

func (r *UserRepo) FindByOrganization(organizationID string) ([]models.User, error) {
    var usersInOrg []models.User
    for _, user := range *r.users {
        if user.OrganizationID == organizationID {
            usersInOrg = append(usersInOrg, user)
        }
    }

    if len(usersInOrg) == 0 {
        // Assuming errs.NewNotFound is a correct call to a custom error handling function
        return nil, errs.NewNotFound("users for organization")
    }

    return usersInOrg, nil
}