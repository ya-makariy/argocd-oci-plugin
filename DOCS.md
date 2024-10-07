## Documentaion
### Instalation
```bash
kubectl apply -k ./manifests/cmp-sidecar -n argocd
kubectl -n argocd create secret generic argocd-oci-plugin-credentials \
--from-literal=AOP_USERNAME='<your_username>' \
--from-file=AOP_PASSWORD=<path/to/your/password> # or --from-literal=AOP_PASSWORD='<your_password>'
kubectl -n argocd patch deployments.apps argocd-repo-server \
-p '{"spec": {"template": {"spec": {"containers": [{"name": "aop-helm", "envFrom":[{"secretRef": {"name": "argocd-oci-plugin-credentials"}}]}, {"name": "aop", "envFrom":[{"secretRef": {"name": "argocd-oci-plugin-credentials"}}]}]}}}}'
```
### Configuration
### Example of usage
Example of application, i prefer use this [one chart](https://github.com/paulbouwer/hello-kubernetes) for testing
```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: hello-kubernetes
  namespace: argocd
spec:
  project: default
  destination:
    server: "https://kubernetes.default.svc"
    namespace: hello-kubernetes
  source:
    repoURL: "<YOUR_REGISTRY>/helm-charts"
    chart: hello-kubernetes
    targetRevision: "1.0.0"
    plugin:
      name: argocd-oci-plugin-helm
      env:
      - name: AOP_REPOSITORY
        value: <YOUR_REGISTRY>/values/hello-kubernetes-values
      - name: AOP_TAG
        value: v1
  syncPolicy:
    syncOptions:
    - CreateNamespace=true
```