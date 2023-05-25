package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermDataFactoryDatasetDelimitedText = `{
  "block": {
    "attributes": {
      "additional_properties": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "annotations": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "column_delimiter": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "compression_codec": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "compression_level": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_factory_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_factory_name": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encoding": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "escape_character": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "first_row_as_header": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "folder": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "linked_service_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "null_value": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parameters": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "quote_character": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "row_delimiter": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "azure_blob_fs_location": {
        "block": {
          "attributes": {
            "file_system": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "filename": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "azure_blob_storage_location": {
        "block": {
          "attributes": {
            "container": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "dynamic_filename_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "dynamic_path_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "filename": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "http_server_location": {
        "block": {
          "attributes": {
            "dynamic_filename_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "dynamic_path_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "filename": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "relative_url": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "schema_column": {
        "block": {
          "attributes": {
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "read": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermDataFactoryDatasetDelimitedTextSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermDataFactoryDatasetDelimitedText), &result)
	return &result
}
