package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azurerm-schema/v2/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAzurermSqlManagedInstanceSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AzurermSqlManagedInstanceSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
