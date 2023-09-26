/*
Delphix DCT API

Testing FeatureFlagApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package delphix_dct_api

import (
	"context"
	"testing"

	openapiclient "github.com/delphix/dct-sdk-go/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_delphix_dct_api_FeatureFlagApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FeatureFlagApiService GetEnabledFeaturesFlag", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.FeatureFlagApi.GetEnabledFeaturesFlag(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
