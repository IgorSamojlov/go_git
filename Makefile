lint:
	golangci-lint run --timeout=2m ./...
tests:
	go test -v
prepare:
	go run . init
	go run . config user.email i.sam
	go run . config user.name i.sam
	echo '.gitignore' >> .git2ignore
	echo '.git2ignore' >> .git2ignore
	echo '.git/*' >> .git2ignore
	echo '.git2/*' >> .git2ignore
	echo '/support' >> .git2ignore
clear:
	rm -rf .git2
setup: clear prepare
