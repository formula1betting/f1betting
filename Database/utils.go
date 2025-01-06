package dbms

import (
	"bufio"
	"os"
	"strings"
)

func loadSQLFile(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	queries := make(map[string]string)
	scanner := bufio.NewScanner(file)
	var currentName string
	var queryBuilder strings.Builder
	inQuery := false

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)

		// Skip empty lines when not in a query
		if !inQuery && trimmedLine == "" {
			continue
		}

		// Handle query name
		if strings.HasPrefix(trimmedLine, "-- name:") {
			// Save previous query if exists
			if currentName != "" && queryBuilder.Len() > 0 {
				queries[currentName] = strings.TrimSpace(queryBuilder.String())
				queryBuilder.Reset()
			}
			// Extract new query name
			currentName = strings.TrimSpace(strings.TrimPrefix(trimmedLine, "-- name:"))
			inQuery = true
			continue
		}

		// Skip comments when in a query
		if inQuery && strings.HasPrefix(trimmedLine, "--") {
			continue
		}

		// Add line to current query
		if inQuery {
			if queryBuilder.Len() > 0 {
				queryBuilder.WriteString("\n")
			}
			queryBuilder.WriteString(line)

			// Check if query ends with semicolon
			if strings.HasSuffix(trimmedLine, ";") {
				queries[currentName] = strings.TrimSpace(queryBuilder.String())
				queryBuilder.Reset()
				inQuery = false
			}
		}
	}

	// Add the last query if exists
	if currentName != "" && queryBuilder.Len() > 0 {
		queries[currentName] = strings.TrimSpace(queryBuilder.String())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return queries, nil
}
