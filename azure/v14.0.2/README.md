# :zap: Giant Swarm Release v14.0.2 for Azure :zap:

This release improves draining of the nodes with fixes on azure-scheduled-events app.

## Change details


### azure-scheduled-events [0.3.0](https://github.com/giantswarm/azure-scheduled-events/releases/tag/v0.3.0)

#### Fixed
- Ensure to wait long enough when draining a node before considering the node drained.
#### Changed
- Change drain timeout to 15 minutes.



