import ctypes

tr = ctypes.CDLL("../library/translator.so")
translator = getattr(tr, "RemoveSysdigLabelsInExpression")
translator.argtypes = [ctypes.c_char_p,ctypes.c_char_p]
translator.restype = ctypes.c_char_p
query = b"sum by (kube_cluster_name, input_namespace, input_container)(rate(fluentd_input_status_num_records_total{kube_cluster_name=~$cluster, input_namespace=~$input_namespace, input_container=~$input_container}[5m])) == 0"
c_query = ctypes.c_char_p(query)
empty = b""
c_empty = ctypes.c_char_p(empty)
returnVale = translator(c_query,c_empty)
print(returnVale)