# Deployment Operations

### Create

The list of fields available to create Deployment:
- Namespace: Required field
- Name: Required field
- ContainerNames: Required field(If more than one containers just pass the name separate by ",". Ex: nginx,apache)
- ContainerImages: Required field(If more than one containers just pass the image separate by ",". Ex: nginx:latest,apache2@latest)
- ContainerPorts: Required filed(If more than one containers just pass the port details separate by "," and if there are more port details for single container separate by "|". Ex: http:80|https:443,http:80)
- Label: Optional field
- Replica: Optional field

### List

The list of fields available to list deployment in particular namespace:
- Namespace: Required field
- Label: Optional field

The list of fields available to list deployment in all namespace:
- Label: Optional field

### Get

The list of fields available to get deployment in particular namespace:
- Namespace: Required field
- Name: Required field

### Delete

The list of fieds available to delete deployment in particular namespace:
- Namespace: Required field
- Name: Required field

### Update

The list of fields available to update deployment in particular namespace(We can update label, annotation, replica or image):
- Namespace: Required field
- Name: Required field
- Label: Optional field
- Annotation: Optional field
- Replica: Optional field
- ContianerName: Optional field
- Image: Optional field(Image is updated based on the contianer name)