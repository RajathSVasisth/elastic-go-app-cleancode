swagger:
	swag init --dir ./cmd/ --output ./cmd/docs --pd --parseInternal --parseDepth 10
generatemocks:
	mockery --all --keeptree --inpackage