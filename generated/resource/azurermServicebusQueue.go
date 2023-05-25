package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermServicebusQueue = `{
  "block": {
    "attributes": {
      "auto_delete_on_idle": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dead_lettering_on_message_expiration": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "default_message_ttl": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "duplicate_detection_history_time_window": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_batched_operations": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_express": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_partitioning": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "forward_dead_lettered_messages_to": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "forward_to": {
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
      "lock_duration": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_delivery_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_message_size_in_kilobytes": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "max_size_in_megabytes": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "namespace_name": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "requires_duplicate_detection": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "requires_session": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "resource_group_name": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
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

func AzurermServicebusQueueSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermServicebusQueue), &result)
	return &result
}
