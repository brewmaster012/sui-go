package suiclient_test

import (
	"encoding/json"
	"testing"

	"github.com/pattonkan/sui-go/sui"
	"github.com/pattonkan/sui-go/suiclient"

	"github.com/stretchr/testify/require"
)

func TestTagJsonUnmarshal(t *testing.T) {
	test := func(str string) suiclient.WrapperTaggedJson[sui.Owner] {
		var s suiclient.WrapperTaggedJson[sui.Owner]
		data := []byte(str)
		err := json.Unmarshal(data, &s)
		require.NoError(t, err)
		return s
	}
	{
		v := test(`"Immutable"`).Data
		require.Nil(t, v.AddressOwner)
		require.Nil(t, v.ObjectOwner)
		require.Nil(t, v.Shared)
		require.NotNil(t, v.Immutable)
	}
	{
		v := test(`{"AddressOwner": "0x7e875ea78ee09f08d72e2676cf84e0f1c8ac61d94fa339cc8e37cace85bebc6e"}`).Data
		require.NotNil(t, v.AddressOwner)
		require.Nil(t, v.ObjectOwner)
		require.Nil(t, v.Shared)
		require.Nil(t, v.Immutable)
	}
}
