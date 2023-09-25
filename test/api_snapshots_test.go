/*
Delphix DCT API

Testing SnapshotsApiService

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

func Test_delphix_dct_api_SnapshotsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SnapshotsApiService CreateSnapshotTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var snapshotId string

		resp, httpRes, err := apiClient.SnapshotsApi.CreateSnapshotTags(context.Background(), snapshotId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SnapshotsApiService DeleteSnapshot", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var snapshotId string

		resp, httpRes, err := apiClient.SnapshotsApi.DeleteSnapshot(context.Background(), snapshotId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SnapshotsApiService DeleteSnapshotTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var snapshotId string

		httpRes, err := apiClient.SnapshotsApi.DeleteSnapshotTags(context.Background(), snapshotId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SnapshotsApiService FindByLocation", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SnapshotsApi.FindByLocation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SnapshotsApiService FindByTimestamp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SnapshotsApi.FindByTimestamp(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SnapshotsApiService GetSnapshotById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var snapshotId string

		resp, httpRes, err := apiClient.SnapshotsApi.GetSnapshotById(context.Background(), snapshotId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SnapshotsApiService GetSnapshotTags", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var snapshotId string

		resp, httpRes, err := apiClient.SnapshotsApi.GetSnapshotTags(context.Background(), snapshotId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SnapshotsApiService GetSnapshotTimeflowRange", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var snapshotId string

		resp, httpRes, err := apiClient.SnapshotsApi.GetSnapshotTimeflowRange(context.Background(), snapshotId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SnapshotsApiService GetSnapshots", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SnapshotsApi.GetSnapshots(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SnapshotsApiService SearchSnapshots", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SnapshotsApi.SearchSnapshots(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SnapshotsApiService UnsetSnapshotRetention", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var snapshotId string

		resp, httpRes, err := apiClient.SnapshotsApi.UnsetSnapshotRetention(context.Background(), snapshotId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SnapshotsApiService UpdateSnapshot", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var snapshotId string

		resp, httpRes, err := apiClient.SnapshotsApi.UpdateSnapshot(context.Background(), snapshotId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}