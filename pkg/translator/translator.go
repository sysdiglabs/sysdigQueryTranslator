package translator

import "regexp"

func RemoveSysdigLabels(expr, excludeScope string) string {
	// Remove scopes of sysdig kube_cluster_name=~$cluster for example
	scope := "kube_cluster_name|kube_namespace_name|kube_pod_name|kube_workload_name|kube_node_name|agent_tag_cluster|instance|node|namespace|host_hostname"
	if excludeScope != "" {
		scope = scope + "|" + excludeScope
	}
	regex := regexp.MustCompile("(?:" + scope + ")" + "(?:=|=~|!=|!~)\\$[\\w_]*,")
	expr = regex.ReplaceAllString(expr, "")
	regex = regexp.MustCompile("(?:" + scope + ")" + "(?:=|=~|!=|!~)\"\",")
	expr = regex.ReplaceAllString(expr, "")
	regex = regexp.MustCompile("(?:" + scope + ")" + "(?:=|=~|!=|!~)\\$[\\w_]*\\}")
	expr = regex.ReplaceAllString(expr, "}")
	regex = regexp.MustCompile("(?:" + scope + ")" + "(?:=|=~|!=|!~)\"\"}")
	expr = regex.ReplaceAllString(expr, "}")
	// Remove the sysdig aggregation
	aggregation := "kube_cluster_name|kube_namespace_name|kube_pod_name|kube_workload_name|kube_node_name|agent_tag_cluster"
	aggregation = aggregation + "|" + excludeScope
	regex = regexp.MustCompile("(?:" + aggregation + ")")
	expr = regex.ReplaceAllString(expr, "")
	// Remove extra commas
	regex = regexp.MustCompile("(?:,+)(?:\\s*)(?:,+)")
	expr = regex.ReplaceAllString(expr, ",")
	regex = regexp.MustCompile(",\\s*\\}")
	expr = regex.ReplaceAllString(expr, "}")
	regex = regexp.MustCompile("\\{\\s*,")
	expr = regex.ReplaceAllString(expr, "{")
	regex = regexp.MustCompile(",\\s*\\)")
	expr = regex.ReplaceAllString(expr, ")")
	regex = regexp.MustCompile("\\(\\s*,")
	expr = regex.ReplaceAllString(expr, "(")
	// Remove duplicate \\
	regex = regexp.MustCompile("\\\\")
	expr = regex.ReplaceAllString(expr, "")
	// Remove empty by ()
	regex = regexp.MustCompile("by\\s*\\(\\s*\\)")
	expr = regex.ReplaceAllString(expr, "")
	// Remove empty scope
	regex = regexp.MustCompile("\\{\\s*\\}")
	expr = regex.ReplaceAllString(expr, "")
	// Remove 2 or more whitespaces
	regex = regexp.MustCompile("[ \\t]{2,}")
	expr = regex.ReplaceAllString(expr, " ")
	// Remove the extra\n
	regex = regexp.MustCompile("\\n")
	expr = regex.ReplaceAllString(expr, "")
	// Add trailing \n, very important don't touch
	expr = expr + "\n"
	return expr
}
