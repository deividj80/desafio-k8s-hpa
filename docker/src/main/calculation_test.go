package main

import (
	"testing"
)

func TestCalculation(t *testing.T)  {
	
	cases := []struct {
		value float64
		result float64
	} {
		{0.22, 0.689041575982343},
		{1.5, 2.724744871391589},
		{5, 7.23606797749979},
	}

	for _, test := range cases {

		t.Logf("Testando caso: %v \n",test.value)
		
		tests := Calculation(test.value)

		if tests != test.result {

			t.Errorf("Correto: %v - Resultado: %v", test.result, tests)

		}
	}
}