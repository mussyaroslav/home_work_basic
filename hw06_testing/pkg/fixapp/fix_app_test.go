package fixapp_test

import (
	"testing"

	"github.com/mussyaroslav/home_work_basic/hw06_testing/pkg/fixapp"
	"github.com/mussyaroslav/home_work_basic/hw06_testing/pkg/fixapp/types"
	"github.com/stretchr/testify/require"
)

func TestFixAppGood(t *testing.T) {
	staff, err := fixapp.Fixapp("data.json")
	req := []types.Employee{
		{UserID: 10, Age: 25, Name: "Rob", DepartmentID: 3},
		{UserID: 11, Age: 30, Name: "George", DepartmentID: 2},
	}
	require.NoError(t, err)
	require.Equal(t, staff, req)
}

func TestFixAppBadPatch(t *testing.T) {
	_, err := fixapp.Fixapp("baddata.json")
	require.Error(t, err)
}

func TestFixAppBadJSON(t *testing.T) {
	_, err := fixapp.Fixapp("baddata.json")
	require.Error(t, err)
}
