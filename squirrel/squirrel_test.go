package squirrel

import (
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/blockloop/scan"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSquirrel_ColumnsValues(t *testing.T) {
	type User struct {
		ID   int
		Name string
		Age  int
	}

	user := &User{
		ID:   1,
		Name: "Brett",
		Age:  100,
	}

	cols, err := scan.Columns(user)
	require.NoError(t, err)

	vals, err := scan.Values(cols, user)
	require.NoError(t, err)

	actualQuery, actualValues, err := sq.Insert("users").
		Columns(cols...).
		Values(vals...).
		ToSql()
	require.NoError(t, err)

	expectedQuery := `INSERT INTO users (ID,Name,Age) VALUES (?,?,?)`

	assert.Equal(t, expectedQuery, actualQuery)
	assert.Equal(t, []interface{}{user.ID, user.Name, user.Age}, actualValues)
}
