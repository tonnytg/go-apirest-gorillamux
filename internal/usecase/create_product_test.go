package usecase

import "testing"

func TestCreateProduct(t *testing.T) {
	c := CreateProductUseCase{}
	_, err := c.Execute(CreateProductInputDto{})
	if err != nil {
		t.Error("Error")
	}
}
