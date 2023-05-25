package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSentinelAlertRuleScheduled = `{
  "block": {
    "attributes": {
      "alert_rule_template_guid": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_analytics_workspace_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "query": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "query_frequency": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "query_period": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "severity": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "suppression_duration": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "suppression_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tactics": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "trigger_operator": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "trigger_threshold": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "event_grouping": {
        "block": {
          "attributes": {
            "aggregation_method": {
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
      "incident_configuration": {
        "block": {
          "attributes": {
            "create_incident": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "block_types": {
            "grouping": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "entity_matching_method": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "group_by": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "lookback_duration": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "reopen_closed_incidents": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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

func AzurermSentinelAlertRuleScheduledSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSentinelAlertRuleScheduled), &result)
	return &result
}
