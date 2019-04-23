package service

import (
	"QianJi_Excel/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchingRuleService_FindTypeMatchingRule(t *testing.T) {
	service := NewMatchingRuleService()
	resultList, err := service.FindTypeMatchingRule(1)
	assert.NoError(t,err)
	assert.NotEmpty(t,resultList)
	utils.PrintStruct(resultList)
}
