docker-build:
	docker buildx build --platform linux/arm/v7 . -t gobug

.PHONY: docker-build
