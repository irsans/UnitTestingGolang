package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"errors"
	"fmt"
)

type CategoryService struct {
	//attribut repository
	Repository repository.CategoryRepository
}

func (service CategoryService) Gets(Id string) (*entity.Category, error)  {
	fmt.Println("Test 0")
	category := service.Repository.FindById(Id)
	fmt.Println("Test Get", category)
	if category.Id != Id{
		return nil , errors.New("ID tidak MATCH")
	}
	if category == nil {
		return nil, errors.New("Category Not Found")
	}else {
		return category, nil
	}
}