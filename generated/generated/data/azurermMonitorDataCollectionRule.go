package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMonitorDataCollectionRule = `{
  "block": {
    "attributes": {
      "data_collection_endpoint_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_flow": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "built_in_transform": "string",
              "destinations": [
                "list",
                "string"
              ],
              "output_stream": "string",
              "streams": [
                "list",
                "string"
              ],
              "transform_kql": "string"
            }
          ]
        ]
      },
      "data_sources": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "data_import": [
                "list",
                [
                  "object",
                  {
                    "event_hub_data_source": [
                      "list",
                      [
                        "object",
                        {
                          "consumer_group": "string",
                          "name": "string",
                          "stream": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "extension": [
                "list",
                [
                  "object",
                  {
                    "extension_json": "string",
                    "extension_name": "string",
                    "input_data_sources": [
                      "list",
                      "string"
                    ],
                    "name": "string",
                    "streams": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "iis_log": [
                "list",
                [
                  "object",
                  {
                    "log_directories": [
                      "list",
                      "string"
                    ],
                    "name": "string",
                    "streams": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "log_file": [
                "list",
                [
                  "object",
                  {
                    "file_patterns": [
                      "list",
                      "string"
                    ],
                    "format": "string",
                    "name": "string",
                    "settings": [
                      "list",
                      [
                        "object",
                        {
                          "text": [
                            "list",
                            [
                              "object",
                              {
                                "record_start_timestamp_format": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "streams": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "performance_counter": [
                "list",
                [
                  "object",
                  {
                    "counter_specifiers": [
                      "list",
                      "string"
                    ],
                    "name": "string",
                    "sampling_frequency_in_seconds": "number",
                    "streams": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "platform_telemetry": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "streams": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "prometheus_forwarder": [
                "list",
                [
                  "object",
                  {
                    "label_include_filter": [
                      "list",
                      [
                        "object",
                        {
                          "label": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "name": "string",
                    "streams": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "syslog": [
                "list",
                [
                  "object",
                  {
                    "facility_names": [
                      "list",
                      "string"
                    ],
                    "log_levels": [
                      "list",
                      "string"
                    ],
                    "name": "string",
                    "streams": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "windows_event_log": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "streams": [
                      "list",
                      "string"
                    ],
                    "x_path_queries": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "windows_firewall_log": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "streams": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "destinations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "azure_monitor_metrics": [
                "list",
                [
                  "object",
                  {
                    "name": "string"
                  }
                ]
              ],
              "event_hub": [
                "list",
                [
                  "object",
                  {
                    "event_hub_id": "string",
                    "name": "string"
                  }
                ]
              ],
              "event_hub_direct": [
                "list",
                [
                  "object",
                  {
                    "event_hub_id": "string",
                    "name": "string"
                  }
                ]
              ],
              "log_analytics": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "workspace_resource_id": "string"
                  }
                ]
              ],
              "monitor_account": [
                "list",
                [
                  "object",
                  {
                    "monitor_account_id": "string",
                    "name": "string"
                  }
                ]
              ],
              "storage_blob": [
                "list",
                [
                  "object",
                  {
                    "container_name": "string",
                    "name": "string",
                    "storage_account_id": "string"
                  }
                ]
              ],
              "storage_blob_direct": [
                "list",
                [
                  "object",
                  {
                    "container_name": "string",
                    "name": "string",
                    "storage_account_id": "string"
                  }
                ]
              ],
              "storage_table_direct": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "storage_account_id": "string",
                    "table_name": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "identity": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "identity_ids": [
                "list",
                "string"
              ],
              "principal_id": "string",
              "tenant_id": "string",
              "type": "string"
            }
          ]
        ]
      },
      "immutable_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "kind": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stream_declaration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "column": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "type": "string"
                  }
                ]
              ],
              "stream_name": "string"
            }
          ]
        ]
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "read": {
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

func AzurermMonitorDataCollectionRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMonitorDataCollectionRule), &result)
	return &result
}
