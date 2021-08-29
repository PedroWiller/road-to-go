// TESTE DE UNIDADES

package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioTest struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoEndereco(t *testing.T) {
	t.Parallel() //indica que ira rodar em paralelo com outros testes

	//Padrao TextXxxxxx
	cenariosTest := []cenarioTest{
		{"Rua ABC", "Rua"},
		{"Avenida ABC", "Avenida"},
		{"Rodovia dos imigrantes", "Rodovia"},
		// {"Praça das rosas", "Tipo Invalido"},
		{"RUA DOS BOBOS", "Rua"},
	}

	for _, cenario := range cenariosTest {
		retorno := TipoEndereco(cenario.enderecoInserido)

		if retorno != cenario.retornoEsperado {
			t.Errorf("O tipo de resposta é diferente do esperado, esperava %s e recebeu %s",
				cenario.retornoEsperado, cenario.enderecoInserido)
		}

	}
}

func TestAdicional(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Errorf("erro para teste")
	}
}
