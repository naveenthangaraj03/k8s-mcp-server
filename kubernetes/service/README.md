# Service Operations

### Create

The list of field available to create service
- Namespace: Required field
- Name: Required field
- SelectorLabel: Required field
- TargetPort: Required field(If more than one target port just separate by ",". Ex: 8080,9090)
- ServicePort: Required field(If more than one service port details just separate by ",". Ex: http:8080,metrics:9090)
- ServiceType: Optional field

### List

The list of fields available to list servcie in particular namespace:
- Namespace: Required field

No fields is required to list service in all namespace.

### Get

The list of fields available to get service in particular namespace:
- Namespace: Required field
- Name: Required field

### Delete

The list of fieds available to delete service in particular namespace:
- Namespace: Required field
- Name: Required field

### Update

The list of fields available to update servcie in particular namespace(We can update only selector label or service type):
- Namespace: Required field
- Name: Required field
- SelectorLabel: Optional field
- Service Type: Optional field