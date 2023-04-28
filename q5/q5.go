package q5

import (
	"strings"
	"unicode"
)

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.

func ProcessString(s string) string {
	strings.ReplaceAll(s, old: "A"new: "")
	strings.ReplaceAll(s, old: "E"new: "")
	strings.ReplaceAll(s, old: "I", new "")
	strings.ReplaceAll(s, old: "O", new "")
	strings.ReplaceAll(s, old: "U", new "")
	strings .ReplaceAll(s, old; "a", new "")
	strings.ReplaceAll(s, old; "e", new: "")
	strings. ReplaceAll(s, old: "i", new"")
	strings.ReplaceAll(s, old: "o", new"")
	strings.ReplaceAll(s, old: "u", new"")

	strings. ToLower (s)

	str := ""

	for 1 := range s {

		if strings.Contains( s "BCDFGHJKLMNPQRSTVWXYZbcdfabiklmnparstvwxyz", string(s[i])) {
		res :="." + strings.ToLower
		str += res
		}
	}
		return str
	return ""
}
