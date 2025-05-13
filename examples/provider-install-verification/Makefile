# List: TRACE, DEBUG, INFO, WARN, ERROR
LOG_LEVEL	=	DEBUG

re: destroy apply

plan:
	terraform plan -var-file=tester.tfvars
apply:
	TF_LOG=$(LOG_LEVEL) TF_LOG_PATH=trace.txt terraform apply -auto-approve -var-file=tester.tfvars
output:
	terraform output -json manta_version | jq -r .version
destroy:
	terraform destroy -auto-approve -var-file=tester.tfvars
