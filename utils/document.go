package utils

import "strings"

func CleanCPF(cpf string) string {
	cpf = strings.ReplaceAll(cpf, "-", "")
	cpf = strings.ReplaceAll(cpf, ".", "")
	return cpf
}
