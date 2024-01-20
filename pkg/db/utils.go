package db

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

func replaceDbName(dbUrl, dbName string) string {
	parsed, err := url.Parse(dbUrl)
	if err != nil {
		return dbUrl
	}

	parsed.Path = "/" + dbName

	return parsed.String()
}

func parseDatabaseName(dbUrl string) (string, error) {
	parsed, err := url.Parse(dbUrl)
	if err != nil {
		return "", err
	}

	path := strings.TrimPrefix(parsed.Path, "/")

	if path == "" {
		return "", fmt.Errorf("empty db name")
	}

	return path, nil
}

func makeMigrateUrl(dbUrl string) string {
	urlRe := regexp.MustCompile("^[^\\?]+")
	url := urlRe.FindString(dbUrl)

	sslModeRe := regexp.MustCompile("(sslmode=)[a-zA-Z0-9]+")
	sslMode := sslModeRe.FindString(dbUrl)

	return fmt.Sprintf("%s?%s", url, sslMode)
}

func FormatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", "")
}
