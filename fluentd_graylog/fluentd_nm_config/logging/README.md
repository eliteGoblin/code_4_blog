
```
kubectl apply -f fluentd-rbac.yaml
ktmpl fluentd-log-collector.yaml --parameter-file clusters/$c.yaml |  kubectl apply -f -
```
