# Pod Operations

### Create

The list of field available to create pod:
- Namespace: Required field
- Name: Required field
- ContainerNames: Required field(If more than one containers just pass the name separate by ",". Ex: nginx,apache)
- ContainerImages: Required field(If more than one containers just pass the image separate by ",". Ex: nginx:latest,apache2@latest)
- ContainerPorts: Required filed(If more than one containers just pass the port details separate by "," and if there are more port details for single container separate by "|". Ex: http:80|https:443,http:80)
- Label: Optional field

### List

The list of fields available to list pods in particular namespace:
- Namespace: Required field
- Label: Optional field

The list of fields available to list pods in all namespace:
- Label: Optional field

### Get

The list of fields available to get pod in particular namespace:
- Namespace: Required field
- Name: Required field

### Delete

The list of fieds available to delete pods in particular namespace:
- Namespace: Required field
- Name: Required field

### Update

The list of fields available to update pods in particular namespace(We can update only label):
- Namespace: Required field
- Name: Required field
- label: Required field

### Logs

The list of fields available for pod logs:
- Namespace: Required field
- Name: Required field
- ContainerName: Required field
- Tailline: Optional field
