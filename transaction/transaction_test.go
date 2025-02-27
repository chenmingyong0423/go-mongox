package transaction

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"testing"
)

type Student struct {
	Name string
}

func TestTx_Auto(t *testing.T) {
	tx := NewTransaction[Student](&mongo.Collection{})
	assert.NotNil(t, tx)
}
