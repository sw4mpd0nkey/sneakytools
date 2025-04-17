package injectors

import "fmt"

type LdapInjector struct {
	Url      string
	Username string
}

func NewLdapInjector(u, n string) *LdapInjector {
	return &LdapInjector{
		Url:      u,
		Username: n,
	}
}

func (li *LdapInjector) TestPassword() {
	fmt.Println("TestPassword")
}

func RunInject() {
	fmt.Println("stiny")
}
