/*
Delphix DCT API

Testing ConnectorsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package delphix_dct_api

import (
	"context"
	"testing"

	openapiclient "github.com/delphix/dct-sdk-go/v14"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_delphix_dct_api_ConnectorsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ConnectorsApiService ConnectorsTest", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connectorId string

		resp, httpRes, err := apiClient.ConnectorsApi.ConnectorsTest(context.Background(), connectorId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsApiService CreateConnectorTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connectorId string

		resp, httpRes, err := apiClient.ConnectorsApi.CreateConnectorTags(context.Background(), connectorId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsApiService DeleteConnectorTag", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connectorId string

		httpRes, err := apiClient.ConnectorsApi.DeleteConnectorTag(context.Background(), connectorId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsApiService GetConnectorById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connectorId string

		resp, httpRes, err := apiClient.ConnectorsApi.GetConnectorById(context.Background(), connectorId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsApiService GetConnectorTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connectorId string

		resp, httpRes, err := apiClient.ConnectorsApi.GetConnectorTags(context.Background(), connectorId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsApiService GetConnectors", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ConnectorsApi.GetConnectors(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsApiService SearchConnectors", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ConnectorsApi.SearchConnectors(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsApiService UpdateConnectorById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connectorId string

		resp, httpRes, err := apiClient.ConnectorsApi.UpdateConnectorById(context.Background(), connectorId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
