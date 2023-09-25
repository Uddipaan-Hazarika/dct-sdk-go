/*
Delphix DCT API

Testing DatabaseTemplatesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package delphix_dct_api

import (
	"context"
	"testing"

	openapiclient "github.com/delphix/dct-sdk-go/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_delphix_dct_api_DatabaseTemplatesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DatabaseTemplatesApiService CreateDatabaseTemplate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DatabaseTemplatesApi.CreateDatabaseTemplate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabaseTemplatesApiService CreateDatabaseTemplateTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var databaseTemplateId string

		resp, httpRes, err := apiClient.DatabaseTemplatesApi.CreateDatabaseTemplateTags(context.Background(), databaseTemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabaseTemplatesApiService DeleteDatabaseTemplate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var databaseTemplateId string

		resp, httpRes, err := apiClient.DatabaseTemplatesApi.DeleteDatabaseTemplate(context.Background(), databaseTemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabaseTemplatesApiService DeleteDatabaseTemplateTag", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var databaseTemplateId string

		httpRes, err := apiClient.DatabaseTemplatesApi.DeleteDatabaseTemplateTag(context.Background(), databaseTemplateId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabaseTemplatesApiService GetDatabaseTemplateById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var databaseTemplateId string

		resp, httpRes, err := apiClient.DatabaseTemplatesApi.GetDatabaseTemplateById(context.Background(), databaseTemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabaseTemplatesApiService GetDatabaseTemplateTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var databaseTemplateId string

		resp, httpRes, err := apiClient.DatabaseTemplatesApi.GetDatabaseTemplateTags(context.Background(), databaseTemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabaseTemplatesApiService GetDatabaseTemplates", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DatabaseTemplatesApi.GetDatabaseTemplates(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabaseTemplatesApiService ImportDatabaseTemplates", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DatabaseTemplatesApi.ImportDatabaseTemplates(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabaseTemplatesApiService SearchDatabaseTemplates", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DatabaseTemplatesApi.SearchDatabaseTemplates(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabaseTemplatesApiService UndoImportDatabaseTemplates", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DatabaseTemplatesApi.UndoImportDatabaseTemplates(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabaseTemplatesApiService UpdateDatabaseTemplate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var databaseTemplateId string

		resp, httpRes, err := apiClient.DatabaseTemplatesApi.UpdateDatabaseTemplate(context.Background(), databaseTemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
