package githooks

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestCheckAuthorEmail(t *testing.T) {
	NewWithT(t).Expect(CheckAuthorEmail("xxx@qq.com")).NotTo(BeNil())
	NewWithT(t).Expect(CheckAuthorEmail("xxx@rockontrol.com")).To(BeNil())
}
