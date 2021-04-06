package repository

import (
	"belajar-golang-unit-test/entity"
	"fmt"
	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	fmt.Println("Test 1")

	arguments := repository.Mock.Called(id)
	fmt.Println("argument",arguments)

if arguments.Get(0) == nil {
	fmt.Println("Test argument",arguments.Get(0))
	return nil
}else {
	category := arguments.Get(0).(entity.Category)
	fmt.Println("Test arguments",arguments.Get(0))

	return  &category
}
}
