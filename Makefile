.PHONY: run
run:
	PRODUCTION=true HOSTS="host1:host2:host3" DURATION=1s go run main.go

# panic: env: required environment variable "PRODUCTION" is not set
.PHONY: panic
panic:
	 HOSTS="host1:host2:host3" DURATION=1s go run main.go
