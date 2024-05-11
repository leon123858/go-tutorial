all:
	echo "this is a sample for k8s"

node:
	kubectl get nodes

apply:
	kubectl apply -f ./k8s.yaml
	echo "preview on: http://127.0.0.1/swagger/index.html"
delete:
	kubectl delete -f ./k8s.yaml -R
	kubectl delete persistentvolumeclaim redis-storage-redis-0
	kubectl delete persistentvolumeclaim postgres-storage-postgres-0

docker:
	#docker compose up -d
	docker compose down

origin:
	echo "preview on: https://leon1234858.notion.site/GKE-Cloud-SQL-Memorystore-08d4a16f0d2544f0a330af8ea49251a9?pvs=4"
