lint:
	golangci-lint run --timeout=2m ./...
tests:
	go test -v
prepare:
	go run . init
	go run . config user.email i.sam
	go run . config user.name i.sam
	echo '.git?' >> .git2ignore
	echo '.git' >> .git2ignore
clear:
	rm -rf .git2
