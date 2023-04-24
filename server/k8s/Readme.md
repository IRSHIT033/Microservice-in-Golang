### Deploy with kubernates manifests

1. Apply and checking configmap for postgres

```
kubectl apply -f postgres-config.yaml
kubectl get configmap 
```

2. Apply and checking Persistent Volume Claim for postgres
```
kubectl apply -f postgres-pvc-pv.yaml
kubectl get pvc
```

3. Apply and checking Deployment for postgres
```
kubectl apply -f postgres-deployment.yaml
kubectl get deployments
```

4. 

```
kubectl apply -f postgres-service.yaml
kubectl get all
```