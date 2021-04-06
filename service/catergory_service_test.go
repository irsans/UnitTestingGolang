package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

/*
func TestCategoryService_GetNotFound(t *testing.T) {
	//program mock
	categoryRepository.Mock.On("FindById","1").Return(nil)
	category,err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}
*/


func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{

		Id:   "1",
		Name: "Handphone",
	}



	categoryRepository.Mock.On("FindById", "2").Return(category)


	result, err := categoryService.Gets("2")
	assert.Nil(t, err, "Hasil Nil")
	assert.NotNil(t, result, "Hasil Tidak Nil")
	assert.Equal(t, category.Name , result.Name, "Hasil Nama Sama")
	//fmt.Println(err)
	fmt.Println("Test 1",result,err)


}