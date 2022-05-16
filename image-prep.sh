# tag existing image
docker tag compliance-service localhost:32000/compliance-service

# push it to local microk8s registry (localhost:32000)
docker push localhost:32000/compliance-service\

# delete existing deployment
microk8s kubectl delete deployment compliance-service


# create new deployment
microk8s kubectl create deployment compliance-service --image=localhost:32000/compliance-service:latest


# expose deployment 
microk8s kubectl expose deployment compliance-service --type=NodePort --port=8092 --name=compliance-expose

#display list of services from node: 
 microk8s kubectl get all --all-namespaces



