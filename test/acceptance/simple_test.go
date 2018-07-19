package acceptance

import (
	"io/ioutil"
	"net/http"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Acceptance", func() {
	const timeout = 1

	It("Should build something", func() {
		command := exec.Command(pathSimpleApi)
		session, _ := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).ToNot(HaveOccurred())
		Eventually(func() string { return string(session.Out.Contents()) }, timeout).Should(ContainSubstring("Started"))
	})

	It("Should listen on Port 8090", func() {
		resp, err := http.Get("http://localhost:8090")
		Expect(err).ToNot(HaveOccurred())
		Eventually(func() string {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			return string(bodyBytes)
		}, timeout).Should(ContainSubstring("I'm an API"))
	})

	It("Should return a error resp code for an invalid url", func() {
		resp, err := http.Get("http://localhost:8090/stupid")
		Expect(err).ToNot(HaveOccurred())
		Eventually(func() string {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			return string(bodyBytes)
		}, timeout).ShouldNot(ContainSubstring("different"))

	})

	It("Should return something different for api ", func() {
		resp, err := http.Get("http://localhost:8090/api")
		Expect(err).ToNot(HaveOccurred())
		Eventually(func() int {
			return resp.StatusCode
		}, timeout).Should(ContainSubstring("different"))

	})

	It("Should do something different for a different path", func() {
		resp, err := http.Get("http://localhost:8090/api")
		Expect(err).ToNot(HaveOccurred())
		Eventually(func() string {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			return string(bodyBytes)
		}, timeout).Should(ContainSubstring("different"))

	})
})
