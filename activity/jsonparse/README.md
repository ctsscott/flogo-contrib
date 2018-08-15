# JSON Parse
This activity can be used to parse a json string for a given path.  

This activity does depend on the gabs utility from https://github.com/Jeffail/gabs.

Please run a ```go get github.com/Jeffail/gabs``` to install.

## Installation
### Flogo CLI
```bash
flogo install github.com/ctsscott/flogo-contrib/activity/jsonparse
```

### Flogo Web
```bash
https://github.com/ctsscott/flogo-contrib/activity/jsonparse
```

## Schema
Inputs and Outputs:

```json
{
  "inputs": [
    {
      "name": "jsonPath",
      "type": "string"
    },
    { "name": "jsonString", 
      "type": "string" }
  ],
  "outputs": [
    {
      "name": "output",
      "type": "string"
    }
  ]
}
```
## Settings
| Setting        | Required | Description |
|:---------------|:---------|:------------|
| jsonPath       | True     |             |         
| jsonString     | True     | 