package enderecos

import "strings"

//TipoEndereco
func TipoEndereco(endereco string) string {
	tipoValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	endereco = strings.ToLower(endereco)
	primeiraPalavra := strings.Split(endereco, " ")[0]

	enderecoValido := false
	for _, tipo := range tipoValidos {
		if tipo == primeiraPalavra {
			enderecoValido = true
		}
	}

	if enderecoValido {
		return strings.Title(primeiraPalavra)
	}

	return "Tipo Invalido"
}
