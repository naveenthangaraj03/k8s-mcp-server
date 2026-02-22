# StatefulSet Operations

### Create

The list of fields available to create statefulset(As of now sts creation will support for single container):
- Namespace: Required field
- Name: Required field
- ContainerNames: Optional field
- ContainerImages: Required field
- ContainerPorts: Optional filed
- StorageValue: Required field(Ex: 1Gi)
- MounthPath: Required field
- PVCName: Optional filed
- ServiceType: Optional filed
- ServicePort: Optional filed
- Label: Optional field
- Replica: Optional field

### List

The list of fields available to list statefulset in particular namespace:
- Namespace: Required field
- Label: Optional field

The list of fields available to list statefulset in all namespace:
- Label: Optional field

### Get

The list of fields available to get statefulset in particular namespace:
- Namespace: Required field
- Name: Required field

### Delete

The list of fieds available to delete statefulset in particular namespace:
- Namespace: Required field
- Name: Required field

### Update

The list of fields available to update statefulset in particular namespace(We can update label, annotation, replica or image):
- Namespace: Required field
- Name: Required field
- Label: Optional field
- Annotation: Optional field
- Replica: Optional field
- ContianerName: Optional field
- Image: Optional field(Image is updated based on the contianer name)