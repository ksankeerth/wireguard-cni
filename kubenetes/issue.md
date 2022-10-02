I'm trying to deploy a DaemonSet only on Master Node. The master node has a few labels already and I added a few new labels as well. When I use those labels in DaeomonSet Spec, the k8s not scheduling DS deployment in any nodes.

I have installed K8s using Kubeadm on Ubuntu 20 LTS.

Labels available on each nodes

```yaml

NAME       STATUS   ROLES                  AGE   VERSION   LABELS
kcp        Ready    control-plane,master   24h   v1.24.3   beta.kubernetes.io/arch=amd64,beta.kubernetes.io/os=linux,kubernetes.io/arch=amd64,kubernetes.io/hostname=kcp,kubernetes.io/os=linux,kubernetes.io/test=test,node-role.kubernetes.io/control-plane=true,node-role.kubernetes.io/master=true,node.kubernetes.io/exclude-from-external-load-balancers=,test1=1

````

The master node contains the below labels.

- beta.kubernetes.io/arch=amd64
- beta.kubernetes.io/os=linux
- kubernetes.io/arch=amd64
- kubernetes.io/hostname=kcp
- kubernetes.io/os=linux
- kubernetes.io/test=test
- node-role.kubernetes.io/control-plane=true
- node-role.kubernetes.io/master=true
- node.kubernetes.io/exclude-from-external-load-balancers=
- test1=1

The below is my DaemonSet YAML

```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: wgcni-cluster-manager
  namespace: kube-system
spec:
  selector:
    matchLabels:
      name: fluentd-elasticsearch
  template:
    metadata:
      labels:
        name: fluentd-elasticsearch
    spec:
      nodeSelector:
        kubernetes.io/test: "test"
      containers:
      - name: fluentd-elasticsearch
        image: quay.io/fluentd_elasticsearch/fluentd:v2.5.2
        resources:
          limits:
            memory: 200Mi
          requests:
            cpu: 100m
            memory: 200Mi
        volumeMounts:
        - name: varlog
          mountPath: /var/log
      terminationGracePeriodSeconds: 30
      volumes:
      - name: varlog
        hostPath:
          path: /var/log
```


No DaemonSet for `wgcni-cluster-manager` scheduled on any nodes.

```yaml
sankeerthan@Sankeerthans-MacBook-Pro ~ % kubectl get ds -n kube-system
NAME                    DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR             AGE
cilium                  2         2         2       2            2           kubernetes.io/os=linux    24h
kube-proxy              2         2         2       2            2           kubernetes.io/os=linux    24h
wgcni-cluster-manager   0         0         0       0            0           kubernetes.io/test=test   22h
sankeerthan@Sankeerthans-MacBook-Pro ~ %
```

If I use `kubernetes.io/os=linux` for `NodeSelector`, then I can see `wgcni-cluster-manager` deployed on all the nodes.

I changed the `NameSpace` to `default` and with different `lables`. The observations was the same.