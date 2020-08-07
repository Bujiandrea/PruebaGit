package calculator_test

import (
	"testing"
	"github.com/Gyga8K/golang_tutorials/writing_unit_tests"
)

func TestAdd(t *testing.T) {
	sum := calculator.Add(1, 2)

	if sum != 3 {
		t.Error("La suma no es correcta")
		t.Fail()
	} else {
		t.Log("Test Add finalizado correctamente")
	}
}
