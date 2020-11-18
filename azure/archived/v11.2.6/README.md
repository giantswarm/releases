## :zap:  Tenant Cluster Release 11.2.6 for Azure :zap:

This release adds two instance types to the list of supported instance types and fixes an error that was interrupting cluster creation.

### azure-operator v3.0.7
- Add new instance types: ```Standard_E8a_v4``` and ```Standard_E8as_v4```
- Fix for outdated error matching that was preventing clusters from being bootstrapped.
- Reduce number of Azure API calls when creating, updating and scaling clusters which lowers the risk of exceeding Azure API rate limits and hitting error 429.
