# Example of GO REST API deploy to SAP cloud foundry

Example hello world REST API made using the gin framework ready to be deployed on cloud foundry.
## Deployment
Launch this command from your terminal after you are logged in cloud foundry space
```console
$ cf push golang-api-example -b https://github.com/cloudfoundry/go-buildpack.git 
```
