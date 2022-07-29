package translator

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	Context("When processing a text", func() {
		It("Removes the Sysdig labels in expression", func() {
			originalText := "sum by(kube_cluster_name,kube_namespace_name,kube_workload_name)\n(nginx_up{kube_cluster_name=~$cluster,kube_namespace_name  =~$namespace,kube_workload_name=\"delete-me\", kube_pod_name!=\"\"}) == 0\n"
			expectedText := "sum (nginx_up) == 0"
			processedText := RemoveSysdigLabels(originalText, "")
			Expect(expectedText).To(Equal(processedText))
		})
	})
})
