# Deploy


From the root directory, run
```
cd deployment
helm dependency build

helm install library . -f values.yaml
```
