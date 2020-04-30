### :zap: Dear Giant Swarm customer :zap:

Prior to upgrading to version 11.3.0 replacing CoreOS Container Linux with Flatcar Container Linux, there is a need to accept the legal terms of the image deployment on Azure.

This has to be done for every subscription you are running, otherwise during upgrade or creation of new cluster the deployments will fail with Bad Request of: ```MarketplacePurchaseEligibilityFailed```.


Instructions are passed by Microsoft and they are following:

> To accept legal terms using PowerShell, please use Get-AzureRmMarketplaceTerms and Set-AzureRmMarketplaceTerms API(https://go.microsoft.com/fwlink/?linkid=862451), 
to accept the terms using Azure CLI, please use az vm image accpet-terms (https://go.microsoft.com/fwlink/?linkid=2110637) or deploy via the Azure portal to accept the terms
>

Here you can find data which will be needed for the acceptance of the license:
```
Offer:     "flatcar-container-linux"
Publisher: "kinvolk"
SKU:       "stable"
Version:   "distroVersion"
