run: n-makefile apex n-express

n-makefile:
	source .env; go run main.go --id financial-times/n-makefile | xargs echo n-makefile:

apex:
	source .env; go run main.go --id apex/apex | xargs echo apex:

n-express:
	source .env; go run main.go --id financial-times/n-express | xargs echo n-express:
