package services

import (
	"strings"

	"github.com/xwb1989/sqlparser"
)

// Verifica se a string pode ser uma query SQL.
func CanBeSQLQuery(s string) bool {
	keywords := []string{
		"SELECT", "INSERT", "UPDATE", "DELETE", "FROM", "WHERE", "JOIN", "GROUP BY", "ORDER BY", "HAVING",
		"CREATE", "ALTER", "DROP", "TABLE", "INDEX", "DATABASE", "VIEW", "TRIGGER", "PROCEDURE", "FUNCTION",
	}

	upperStr := strings.ToUpper(s)

	for _, keyword := range keywords {
		if strings.Contains(upperStr, keyword) {
			return true
		}
	}

	return false
}

// Verifica se a string é uma query SQL válida.
func IsSQLQuery(s string) bool {
	_, err := sqlparser.Parse(s)
	return err == nil
}
