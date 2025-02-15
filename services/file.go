package services

import (
	"os"
	"path/filepath"
	"strings"
)

// Verifica se o arquivo informado é um arquivo válido para realizar a migração.
func IsValidFile(fileName string) bool {
	return isFileName(fileName) && fileExists(fileName) && isSqlFile(fileName)
}

// Verifica se a string informado é um possível arquivo.
func isFileName(fileName string) bool {
	base := filepath.Base(fileName)

	return base != "" && base != "." && base != "/"
}

// Verifica se o arquivo existe.
func fileExists(fileName string) bool {
	info, err := os.Stat(fileName)

	return !os.IsNotExist(err) && !info.IsDir()
}

// Verifica se o arquivo se trata de um arquivo .sql
func isSqlFile(fileName string) bool {
	return strings.HasSuffix(fileName, ".sql")
}
