package services

import "strings"

func IsSQLQuery(s string) bool {
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
