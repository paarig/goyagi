package application

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func testNew(t *testing.T) {
  app, err := New()
  assert.NoError(t, err)
  assert.NotNil(t, app)
  assert.NotNil(t, app.Config)
}
