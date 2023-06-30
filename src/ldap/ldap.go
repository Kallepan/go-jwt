package ldap

import (
	"fmt"
	"log"

	"github.com/go-ldap/ldap"
)

type LDAPInfo struct {
	BindUsername string
	BindPassword string
	FQDN         string
	BaseDN       string
	Filter       string
	SSLPort      string
	Port         string
}

var Instance *ldap.Conn

func Connect(lInfo LDAPInfo) error {
	l, err := ldap.DialURL(fmt.Sprintf("ldap://%s:%s", lInfo.FQDN, lInfo.Port))

	if err != nil {
		return err
	}

	log.Println("Connected to LDAP!")

	Instance = l
	return nil
}

func ConnectTLS(lInfo LDAPInfo) error {
	l, err := ldap.DialURL(fmt.Sprintf("ldaps://%s:%s", lInfo.FQDN, lInfo.SSLPort))

	if err != nil {
		return err
	}

	log.Println("Connected to LDAP with TLS!")

	Instance = l
	return nil
}

func AnonymousBind(lInfo LDAPInfo) (*ldap.SearchResult, error) {
	Instance.UnauthenticatedBind("")

	anonReq := ldap.NewSearchRequest(
		"",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		lInfo.Filter,
		[]string{""},
		nil,
	)
	result, err := Instance.Search(anonReq)
	if err != nil {
		return nil, fmt.Errorf("Error searching: %s", err)
	}

	if len(result.Entries) > 1 {
		return nil, fmt.Errorf("Too many entries returned")
	}

	if len(result.Entries) == 0 {
		return nil, fmt.Errorf("No entries returned")
	}

	return result, nil
}

func BindAndSearch(lInfo LDAPInfo) (*ldap.SearchResult, error) {
	err := Instance.Bind(lInfo.BindUsername, lInfo.BindPassword)
	if err != nil {
		return nil, fmt.Errorf("Error binding: %s", err)
	}

	searchReq := ldap.NewSearchRequest(
		lInfo.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		lInfo.Filter,
		[]string{""},
		nil,
	)
	result, err := Instance.Search(searchReq)
	if err != nil {
		return nil, fmt.Errorf("Error searching: %s", err)
	}

	if len(result.Entries) > 1 {
		return nil, fmt.Errorf("Too many entries returned")
	}

	if len(result.Entries) == 0 {
		return nil, fmt.Errorf("No entries returned")
	}

	return result, nil
}
