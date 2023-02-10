package usecase_test

import (
	"github.com/tonnytg/go-apirest/internal/usecase"
	"testing"
)

func TestExecute(t *testing.T) {
	l := usecase.ListProductsUseCase{}
	_, err := l.Execute()
	if err != nil {
		t.Error("Error")
	}
}
