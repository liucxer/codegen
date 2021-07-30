package githooks

import (
	"fmt"
	"strings"
)

var emailDomains = []string{
	"datatom.com",
}

func CheckAuthorEmail(email string) error {
	parts := strings.Split(email, "@")
	if len(parts) == 2 {
		for _, e := range emailDomains {
			if parts[1] == e {
				return nil
			}
		}
	}
	return fmt.Errorf("invalid email %s, domain should be one of %v, please use `git config user.email <email>` to set", email, emailDomains)
}
