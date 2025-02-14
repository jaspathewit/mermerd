package util

import (
	"fmt"
	"strings"
)

const fullstopSeparator = "."

// QuoteTableName quotes the table name if it contains a fullstop separator.
func QuoteTableName(tableName string) string {
	if !strings.Contains(tableName, fullstopSeparator) {
		return tableName
	}
	return fmt.Sprintf("\"%s\"", tableName)
}
