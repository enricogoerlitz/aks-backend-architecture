

```sh
$ terraform init

$ terraform workspace list
$ terraform workspace select <env>
$ terraform workspace new <env>

$ terraform validate
$ terraform plan
$ terraform apply
$ terraform apply -auto-approve
$ terraform apply -var-file=tfvars.$(terraform workspace show).tfvars -auto-approve

$ alias tfplan='terraform plan -var-file=tfvars.$(terraform workspace show).tfvars'
$ alias tfplan='terraform plan -var-file=tfvars.$(terraform workspace show).tfvars -var="db_sa_username=adminuser" -var="db_sa_password=adminpw1!"'
$ alias tfapply='terraform apply -var-file=tfvars.$(terraform workspace show).tfvars -auto-approve'
$ alias tfapply='terraform apply -var-file=tfvars.$(terraform workspace show).tfvars -var="db_sa_username=adminuser" -var="db_sa_password=adminpw1!" -auto-approve'
$ tfapply / tfplan -> will execute alias (only valid within an session)

$ terraform destroy
$ terraform destroy -auto-approve
$ terraform destroy -var-file=tfvars.$(terraform workspace show).tfvars -auto-approve

```
