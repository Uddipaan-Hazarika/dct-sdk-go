/*
Delphix DCT API

Testing MaskingEnvironmentsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package delphix_dct_api

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/delphix/dct-sdk-go"
)

func Test_delphix_dct_api_MaskingEnvironmentsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MaskingEnvironmentsApiService GetMaskingEnvironmentById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var maskingEnvironmentId string

		resp, httpRes, err := apiClient.MaskingEnvironmentsApi.GetMaskingEnvironmentById(context.Background(), maskingEnvironmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MaskingEnvironmentsApiService GetMaskingEnvironments", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MaskingEnvironmentsApi.GetMaskingEnvironments(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MaskingEnvironmentsApiService SearchMaskingEnvironments", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MaskingEnvironmentsApi.SearchMaskingEnvironments(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
