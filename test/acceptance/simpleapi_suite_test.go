package acceptance

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
		"github.com/onsi/gomega/gexec"
	"fmt"
)


func TestSimpleapi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Simpleapi Suite")
}

var (
	pathSimpleApi string
	err           error
)

var _ = BeforeSuite(func() {
	pathSimpleApi, err = gexec.Build("github.com/simpleapi")
	fmt.Println(pathSimpleApi)
	Expect(err).ToNot(HaveOccurred())
})

var _ bool = AfterSuite(func() {
	gexec.Kill()
})