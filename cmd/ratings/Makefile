# Image URL to use all building/pushing image targets
IMG ?= uhziel/ratings

docker: docker-build docker-push

docker-build:
	docker build -t ${IMG} -f Dockerfile ../../

docker-push:
	docker push ${IMG}
