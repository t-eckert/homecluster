NAMESPACE_FILES := $(shell find . -name "namespace.yaml")

quadratic-crab:
	@tailscale ssh thomas@quadratic-crab

namespaces:
	@echo "Deploying namespaces..."
	@for file in $(NAMESPACE_FILES); do \
		echo "Applying $$file..."; \
		kubectl apply -f $$file; \
	done
	@echo "All namespaces deployed!"
