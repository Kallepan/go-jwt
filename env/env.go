package env

import (
	"fmt"
	"os"

	"gitlab.com/kallepan/go-jwt/ldap"
)

type DbInfo struct {
	User     string
	Password string
	DbName   string
	Host     string
	Port     string
}

func GetLDAPInfo() ldap.LDAPInfo {
	envSrcUser := "LDAP_BIND_USERNAME"
	envSrcPass := "LDAP_BIND_PASSWORD"
	envSrcFQDN := "LDAP_FQDN"
	envSrcBaseDN := "LDAP_BASE_DN"
	envSrcFilter := "LDAP_FILTER"
	envSrcPort := "LDAP_PORT"
	envSrcSSLPort := "LDAP_SSL_PORT"

	lInfo := ldap.LDAPInfo{
		BindUsername: GetValueFromEnv(envSrcUser, ""),
		BindPassword: GetValueFromEnv(envSrcPass, ""),
		FQDN:         GetValueFromEnv(envSrcFQDN, ""),
		BaseDN:       GetValueFromEnv(envSrcBaseDN, ""),
		Filter:       GetValueFromEnv(envSrcFilter, ""),
		Port:         GetValueFromEnv(envSrcPort, ""),
		SSLPort:      GetValueFromEnv(envSrcSSLPort, ""),
	}

	return lInfo
}

func GetConnectionString() string {
	dbInfo := getDbInfo()

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbInfo.Host,
		dbInfo.User,
		dbInfo.Password,
		dbInfo.DbName,
		dbInfo.Port,
	)

	return connectionString
}

func getDbInfo() DbInfo {
	envSrcUser := "POSTGRES_USER"
	envSrcPass := "POSTGRES_PASSWORD"
	envSrcDbName := "POSTGRES_DB"
	envSrcPort := "POSTGRES_PORT"
	envSrcHost := "POSTGRES_HOST"

	dbInfo := DbInfo{
		User:     GetValueFromEnv(envSrcUser, "test"),
		Password: GetValueFromEnv(envSrcPass, "test"),
		DbName:   GetValueFromEnv(envSrcDbName, "test"),
		Host:     GetValueFromEnv(envSrcHost, "localhost"),
		Port:     GetValueFromEnv(envSrcPort, "5432"),
	}

	return dbInfo
}

func GetValueFromEnv(envSrc string, defaultValue string) string {
	envValue := os.Getenv(envSrc)
	if envValue == "" {
		return defaultValue
	}

	return envValue
}
