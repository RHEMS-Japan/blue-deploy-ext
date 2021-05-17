package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jtblin/go-ldap-client"
)

var (
	key string
	version string
)

func main() {
	base := os.Getenv("LDAP_Base")
	host := os.Getenv("LDAP_Host")
	user := os.Getenv("LDAP_User")
	pass := os.Getenv("LDAP_Pass")
	doc  := os.Getenv("Fs_Doc")

	if len(os.Args) > 1 {
		if os.Args[1] == "-v" {
			fmt.Println(version)
			os.Exit(0)
		}
		if os.Args[1] == "-u" {
			fmt.Print(user)
			os.Exit(0)
		}
		if os.Args[1] == "-t" {
			fmt.Print(doc)
			os.Exit(0)
		}
	}

	client := &ldap.LDAPClient{
		Base:        base,
		Host:        host,
		Port:        389,
		UseSSL:      false,
		SkipTLS:     true,
		BindDN:      fmt.Sprintf("uid=readonlysuer,ou=People,%s", base),
		UserFilter:  "(uid=%s)",
		GroupFilter: "(memberUid=%s)",
		Attributes:  []string{"givenName", "sn", "mail", "uid"},
	}

	defer client.Close()

	client.ServerName = host

	ok, _, err := client.Authenticate(user, pass)
	if err != nil {
		log.Fatalf("auth error , user %s: error %+v", user, err)
	}
	if !ok {
		log.Fatalf("auth failed for user %s", user)
	}
	if key == "" {
		fmt.Println("key is not setting")
	} else {
		fmt.Print(key)
	}
}
