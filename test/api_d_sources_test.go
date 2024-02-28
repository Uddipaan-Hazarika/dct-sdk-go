/*
Delphix DCT API

Testing DSourcesApiService

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

func Test_delphix_dct_api_DSourcesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DSourcesApiService CreateTagsDsource", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dsourceId string

		resp, httpRes, err := apiClient.DSourcesApi.CreateTagsDsource(context.Background(), dsourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DSourcesApiService DeleteDsource", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DSourcesApi.DeleteDsource(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DSourcesApiService DeleteTagsDsource", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dsourceId string

		httpRes, err := apiClient.DSourcesApi.DeleteTagsDsource(context.Background(), dsourceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DSourcesApiService GetAppdataDsourceLinkingDefaults", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DSourcesApi.GetAppdataDsourceLinkingDefaults(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DSourcesApiService GetAseDsourceLinkingDefaults", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DSourcesApi.GetAseDsourceLinkingDefaults(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DSourcesApiService GetDsourceById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dsourceId string

		resp, httpRes, err := apiClient.DSourcesApi.GetDsourceById(context.Background(), dsourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DSourcesApiService GetDsourceSnapshots", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dsourceId string

		resp, httpRes, err := apiClient.DSourcesApi.GetDsourceSnapshots(context.Background(), dsourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DSourcesApiService GetDsources", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DSourcesApi.GetDsources(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DSourcesApiService GetOracleDsourceLinkingDefaults", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DSourcesApi.GetOracleDsourceLinkingDefaults(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DSourcesApiService GetTagsDsource", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dsourceId string

		resp, httpRes, err := apiClient.DSourcesApi.GetTagsDsource(context.Background(), dsourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DSourcesApiService LinkAppdataDatabase", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DSourcesApi.LinkAppdataDatabase(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DSourcesApiService LinkAseDatabase", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DSourcesApi.LinkAseDatabase(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DSourcesApiService LinkOracleDatabase", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DSourcesApi.LinkOracleDatabase(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DSourcesApiService SearchDsources", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DSourcesApi.SearchDsources(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DSourcesApiService SnapshotDsource", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var dsourceId string

		resp, httpRes, err := apiClient.DSourcesApi.SnapshotDsource(context.Background(), dsourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
