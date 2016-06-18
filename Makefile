run: n-makefile apex n-express

n-makefile:
	@$(env) go run main.go --id financial-times/n-makefile --satisfies '<1.0' | xargs echo n-makefile:

apex:
	@$(env) go run main.go --id apex/apex | xargs echo apex:

n-express:
	@$(env) go run main.go --id financial-times/n-express | xargs echo n-express:

deploy:
	@$(env) apex deploy -s GITHUB_API_KEY=$$GITHUB_API_KEY

deploy-dry:
	@$(env) apex deploy --dry-run -s GITHUB_API_KEY=$$GITHUB_API_KEY

help:
	go run main.go --help

test:
	go test -cover ./...

env = $(shell cat .env | sed 's/^/export /' | tr '\n' ';')
