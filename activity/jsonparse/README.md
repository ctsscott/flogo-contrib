# JSON Parse
This activity can be used to parse a json string for a given path of given type.  Three different input sets exist to get through seperate output values.

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
"inputs": [
    { "name": "jsonPath1","type": "string", "required": true},
    { "name": "jsonString1", "type": "string", "required": true},
    { "name": "jsonType1",
      "type": "string",
      "required": true, 
      "allowed": [
        "string",
        "boolean"
      ]
    },
    { "name": "jsonPath2","type": "string"},
    { "name": "jsonString2", "type": "string"},
    { "name": "jsonType2",
      "type": "string",
      "allowed": [
        "string",
        "boolean"
      ]
    },
    { "name": "jsonPath3","type": "string"},
    { "name": "jsonString3", "type": "string"},
    { "name": "jsonType3",
      "type": "string",
      "allowed": [
        "string",
        "boolean"
      ]
    }
  ],
  "outputs": [
    { "name": "output1","type": "string"},
    { "name": "output2", "type": "string" },
    { "name": "output3", "type": "string" }
  ]
```
## Settings
| Setting        | Required | Description |
|:---------------|:---------|:------------|
| jsonPath1      | True     |             |         
| jsonString1    | True     |             |
| jsonType1      | True     |             |
| jsonPath2      | False    |             |         
| jsonString2    | False    |             |  
| jsonType2      | False    |             |  
| jsonPath3      | False    |             |       
| jsonString3    | False    |             |  
| jsonType3      | False    |             |  