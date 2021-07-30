package githooks

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestCheckCommitMsg(t *testing.T) {
	t.Run("invalid format", func(t *testing.T) {
		NewWithT(t).Expect(CheckCommitMsg(": test")).NotTo(BeNil())
	})

	t.Run("invalid type", func(t *testing.T) {
		NewWithT(t).Expect(CheckCommitMsg("hehe: test")).NotTo(BeNil())
	})

	t.Run("invalid rel", func(t *testing.T) {
		NewWithT(t).Expect(CheckCommitMsg("chore(deps): [IEP-xxx] test")).NotTo(BeNil())
	})

	t.Run("valid header", func(t *testing.T) {
		NewWithT(t).Expect(CheckCommitMsg("chore: test")).To(BeNil())
		NewWithT(t).Expect(CheckCommitMsg("fix(account): [IEP-111] test")).To(BeNil())
	})
}
