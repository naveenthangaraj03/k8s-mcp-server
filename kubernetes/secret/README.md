# Secret Operations

### Create

The list of field available to create secret:
- Namespace: Required field
- Name: Required field
- Data: Required field(If there are more than one data just separate by ",". Ex: password=Passw0rd@123,username=admin)

### List

The list of fields available to list secret in particular namespace:
- Namespace: Required field

No field is required to list configmap in all namespace.

### Get

The list of fields available to get secret in particular namespace:
- Namespace: Required field
- Name: Required field

### Delete

The list of fieds available to delete secret in particular namespace:
- Namespace: Required field
- Name: Required field