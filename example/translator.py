import ctypes

tr = ctypes.CDLL("../library/translator.so")
translator = getattr(tr, "RemoveSysdigLabelsInExpression")
translator.argtypes = [ctypes.c_char_p,ctypes.c_char_p]
translator.restype = ctypes.c_char_p
query = 'sum by(kube_cluster_name,kube_namespace_name,kube_workload_name)(nginx_up{kube_cluster_name=~$cluster,kube_namespace_name  =~$namespace,kube_workload_name="delete-me", kube_pod_name!=""}) == 0'
c_query = ctypes.c_char_p(bytes(query, "ascii"))
empty = b""
c_empty = ctypes.c_char_p(empty)
returnVale = translator(c_query,c_empty)
print(returnVale)