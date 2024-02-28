/*
Delphix DCT API

Testing ManagementApiService

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

func Test_delphix_dct_api_ManagementApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ManagementApiService CreateEngineTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var engineId string

		resp, httpRes, err := apiClient.ManagementApi.CreateEngineTags(context.Background(), engineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService CreateHashicorpVault", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ManagementApi.CreateHashicorpVault(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService CreateHashicorpVaultTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var vaultId int64

		resp, httpRes, err := apiClient.ManagementApi.CreateHashicorpVaultTags(context.Background(), vaultId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService DeleteEngineTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var engineId string

		httpRes, err := apiClient.ManagementApi.DeleteEngineTags(context.Background(), engineId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService DeleteHashicorpVault", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var vaultId int64

		httpRes, err := apiClient.ManagementApi.DeleteHashicorpVault(context.Background(), vaultId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService DeleteHashicorpVaultTag", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var vaultId int64

		httpRes, err := apiClient.ManagementApi.DeleteHashicorpVaultTag(context.Background(), vaultId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService GetApiClassification", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ManagementApi.GetApiClassification(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService GetEngineTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var engineId string

		resp, httpRes, err := apiClient.ManagementApi.GetEngineTags(context.Background(), engineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService GetHashicorpVault", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var vaultId int64

		resp, httpRes, err := apiClient.ManagementApi.GetHashicorpVault(context.Background(), vaultId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService GetHashicorpVaultTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var vaultId int64

		resp, httpRes, err := apiClient.ManagementApi.GetHashicorpVaultTags(context.Background(), vaultId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService GetHashicorpVaults", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ManagementApi.GetHashicorpVaults(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService GetLdapConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ManagementApi.GetLdapConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService GetMetadataDatabase", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ManagementApi.GetMetadataDatabase(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService GetProxyConfiguration", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ManagementApi.GetProxyConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService GetRegisteredEngine", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var engineId string

		resp, httpRes, err := apiClient.ManagementApi.GetRegisteredEngine(context.Background(), engineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService GetRegisteredEngines", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ManagementApi.GetRegisteredEngines(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService GetSamlConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ManagementApi.GetSamlConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService GetSmtpConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ManagementApi.GetSmtpConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService ListProperties", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ManagementApi.ListProperties(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService RegisterEngine", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ManagementApi.RegisterEngine(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService SearchEngines", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ManagementApi.SearchEngines(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService SearchHashicorpVaults", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ManagementApi.SearchHashicorpVaults(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService UnregisterEngine", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var engineId string

		resp, httpRes, err := apiClient.ManagementApi.UnregisterEngine(context.Background(), engineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService UpdateApiClassification", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ManagementApi.UpdateApiClassification(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService UpdateLdapConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ManagementApi.UpdateLdapConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService UpdateProperties", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ManagementApi.UpdateProperties(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService UpdateProxyConfiguration", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ManagementApi.UpdateProxyConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService UpdateRegisteredEngine", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var engineId string

		resp, httpRes, err := apiClient.ManagementApi.UpdateRegisteredEngine(context.Background(), engineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService UpdateSamlConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ManagementApi.UpdateSamlConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService UpdateSmtpConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ManagementApi.UpdateSmtpConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService ValidateJavaPath", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var engineId string

		httpRes, err := apiClient.ManagementApi.ValidateJavaPath(context.Background(), engineId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService ValidateLdapConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ManagementApi.ValidateLdapConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagementApiService ValidateSmtpConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ManagementApi.ValidateSmtpConfig(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
