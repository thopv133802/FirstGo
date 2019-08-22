package repositories

import (
	"GoFirst/src/domain/models"
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type CommonRepository interface {
	GetAllProducts() ([]models.Product, error)
	GetPromos() ([]models.Product, error)
	GetCustomerByName(string, string) (models.Customer, error)
	GetCustomerByID(int) (models.Customer, error)
	GetProduct(int) (models.Product, error)
	AddUser(models.Customer) (models.Customer, error)
	SignInUser(username, password string) (models.Customer, error)
	SignOutUserById(int) error
	GetCustomerOrdersByID(int) ([]models.Order, error)
}

type CommonRepositoryImpl struct{
	*gorm.DB
}

func (repository CommonRepositoryImpl) GetAllProducts() (products []models.Product, err error) {
	return products, repository.Find(&products).Error
}

func (repository CommonRepositoryImpl) GetPromos() (products []models.Product, err error) {
	return products, repository.
		Where("promotion IS NOT NULL").
		Find(&products).Error
}

func (repository CommonRepositoryImpl) GetCustomerByName(firstname string, lastname string) (customer models.Customer, err error) {
	return customer, repository.
		Where(&models.Customer {FirstName: firstname, LastName: lastname}).
		Find(&customer).Error
}

func (repository CommonRepositoryImpl) GetCustomerByID(id int) (customer models.Customer, err error) {
	return customer, repository.
		Where(&customer, id).
		Error
}

func (repository CommonRepositoryImpl) GetProduct(id int) (product models.Product, err error) {
	return product, repository.First(&product, id).Error
}

func (repository CommonRepositoryImpl) AddUser(customer models.Customer) (models.Customer, error) {
	err := hashPassword(&customer.Password)
	if err != nil {
		return nil, err
	}
	customer.LoggedIn = true
	err = repository.Create(&customer).Error
	customer.Password = ""
	return customer, err
}

func (repository CommonRepositoryImpl) SignInUser(email, password string) (customer models.Customer, err error) {
	result := repository.Table("customers").
		Where(&models.Customer {Email: email})
	err = result.First(&customer).Error
	if err != nil {
		return customer, err
	}
	if !checkPassword(customer.Password, password) {
		return customer, ERROR_INVALID_PASSWORD
	}
	customer.Password = ""
	err = result.Update("loggedin", 1).Error
	if err != nil {
		return customer, err
	}
	return customer, result.Find(&customer).Error
}

func (repository CommonRepositoryImpl) SignOutUserById(id int) error {
	customer := models.Customer {
		Model: gorm.Model {
			ID: uint(id),
		},
	}
	return repository.Table("customers").
		Where(&customer).
		Update("loggedin", 0).
		Error
}

func (repository CommonRepositoryImpl) GetCustomerOrdersByID(id int) (orders []models.Order, err error) {
	return orders, repository.Table("orders").
		Select("*").
		Joins("join customers on customers.id = customer_id").
		Joins("join products on products.id = product_id").
		Where("customer_id = ?", id).
		Scan(&orders).
		Error
}

var ERROR_INVALID_PASSWORD = errors.New("Invalid password")

func hashPassword(password *string) error {
	if password == nil {
		return errors.New("Reference provided for hashing password is nil")
	}
	passwordBytes := []byte(*password)

	//Obtain hashed password via the GenerateFromPassword() method
	hashedBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	//update password string with the hashed version
	*password = string(hashedBytes[:])
	return nil
}

func checkPassword(existingHash, incomingPass string) bool {
	return bcrypt.CompareHashAndPassword(
		[]byte(existingHash), []byte(incomingPass)) == nil
}