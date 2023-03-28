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
		It("Remove all scopes", func() {
			originalText := "irate(elasticsearch_indices_store_throttle_time_seconds_total{kube_cluster_name=$k8s_cluster,cluster=~$es_cluster, host=~$node,kube_namespace_name=~$namespace, kube_workload_name=~$workload}[$__interval]) "
			expectedText := "irate(elasticsearch_indices_store_throttle_time_seconds_total{  }[$__interval]) "
			processedText := RemoveAllScopes(originalText, "all")
			Expect(expectedText).To(Equal(processedText))
		})
		It("Removes all labels", func() {
			originalText := "irate(elasticsearch_indices_store_throttle_time_seconds_total{kube_cluster_name=$k8s_cluster,cluster=~$es_cluster, host=~$node,kube_namespace_name=~$namespace, kube_workload_name=~$workload}[$__interval]) "
			expectedText := "irate(elasticsearch_indices_store_throttle_time_seconds_total[5m]) "
			processedText := RemoveSysdigLabels(originalText, "all")
			Expect(expectedText).To(Equal(processedText))
		})
		It("Remove $__scope", func() {
			originalText := "irate(elasticsearch_indices_store_throttle_time_seconds_total{$__scope, kube_cluster_name=$k8s_cluster}[$__interval])"
			expectedText := "irate(elasticsearch_indices_store_throttle_time_seconds_total[5m]) "
			processedText := RemoveSysdigLabels(originalText, "")
			Expect(expectedText).To(Equal(processedText))
		})
	})
})
