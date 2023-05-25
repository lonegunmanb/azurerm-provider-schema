package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermKustoCluster = `{
  "block": {
    "attributes": {
      "auto_stop_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "data_ingestion_uri": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "disk_encryption_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "double_encryption_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_auto_stop": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_disk_encryption": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_purge": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_streaming_ingest": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "engine": {
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
      "language_extensions": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "public_network_access_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "purge_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "streaming_ingestion_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "trusted_external_tenants": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "uri": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "zones": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "block_types": {
      "identity": {
        "block": {
          "attributes": {
            "identity_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "principal_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "tenant_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
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
      "optimized_auto_scale": {
        "block": {
          "attributes": {
            "maximum_instances": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "minimum_instances": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "sku": {
        "block": {
          "attributes": {
            "capacity": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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
      },
      "virtual_network_configuration": {
        "block": {
          "attributes": {
            "data_management_public_ip_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "engine_public_ip_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "subnet_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermKustoClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermKustoCluster), &result)
	return &result
}
