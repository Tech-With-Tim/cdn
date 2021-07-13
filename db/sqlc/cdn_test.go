package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Tech-With-Tim/cdn/utils"

	"github.com/stretchr/testify/require"
)

func createRandomAsset(t *testing.T) (CreateAssetParams, int64) {
	arg := CreateAssetParams{
		Mimetype:  "application/octet-stream",
		Name:      utils.RandomString(4) + ".bin", //FileName
		Data:      utils.StrToBinary(utils.RandomString(16), 10),
		Name_2:    utils.RandomString(4), //AssetName
		UrlPath:   utils.RandomString(4),
		CreatorID: 735376244656308274,
	}
	//Context.Background() is to provide empty Context For tests
	assetId, err := testQueries.CreateAsset(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, assetId)
	require.NotZero(t, assetId)
	return arg, assetId
}

func cleanup(t *testing.T, asset CreateAssetParams) {
	deleteArgs := DeleteAssetParams{
		UrlPath:   asset.UrlPath,
		CreatorID: asset.CreatorID,
	}
	err := testQueries.DeleteAsset(context.Background(), deleteArgs)
	require.NoError(t, err)
}

func TestQueries_CreateAsset(t *testing.T) {
	store := NewStore(testDB) //testDb is a global var check cdn_test.go
	// run 6 concurrent transactions
	n := 6
	errors := make(chan error)
	results := make(chan int64)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.CreateAssetFile(context.Background(), CreateAssetParams{
				Mimetype:  "application/octet-stream",
				Name:      utils.RandomString(4) + ".bin", //FileName
				Data:      utils.StrToBinary(utils.RandomString(16), 10),
				Name_2:    utils.RandomString(4), //AssetName
				UrlPath:   utils.RandomString(4),
				CreatorID: 735376244656308274,
			})
			errors <- err
			results <- result

		}()
	}

	// check results
	for i := 0; i < n; i++ {
		err := <-errors
		require.NoError(t, err)
		result := <-results
		require.NotEmpty(t, result)
		require.NotZero(t, result)

		asset, err := store.GetAssetDetailsById(context.Background(), result)
		require.NoError(t, err)
		deleteArgs := DeleteAssetParams{
			UrlPath:   asset.UrlPath,
			CreatorID: asset.CreatorID,
		}
		err = testQueries.DeleteAsset(context.Background(), deleteArgs)
		require.NoError(t, err)

	}
	//_, _ = createRandomAsset(t)
}

func TestQueries_GetAssetDetails(t *testing.T) {
	generatedAsset, assetID := createRandomAsset(t)
	assetTest, err := testQueries.GetAssetDetailsByUrl(context.Background(), generatedAsset.UrlPath)
	require.NoError(t, err)
	require.NotEmpty(t, assetTest)
	require.Equal(t, assetID, assetTest.ID)
	require.Equal(t, generatedAsset.Name_2, assetTest.Name)
	assetTestById, err := testQueries.GetAssetDetailsById(context.Background(), assetID)
	require.NoError(t, err)
	require.NotEmpty(t, assetTestById)
	require.Equal(t, assetTest.Name, assetTestById.Name)
	cleanup(t, generatedAsset)
}

func TestQueries_GetFile(t *testing.T) {
	generatedAsset, _ := createRandomAsset(t)
	fileTest, err := testQueries.GetFile(context.Background(),
		generatedAsset.UrlPath)
	require.NoError(t, err)
	require.NotEmpty(t, fileTest)
	require.Equal(t, generatedAsset.Mimetype, fileTest.Mimetype)
	require.Equal(t, generatedAsset.Data, fileTest.Data)

	cleanup(t, generatedAsset)
}

func TestQueries_ListAssetByCreator(t *testing.T) {
	generatedAsset, assetID := createRandomAsset(t)
	//ExpectedRow := ListAssetByCreatorRow{
	//	ID:      assetID,
	//	Name:    generatedAsset.Name_2,
	//	UrlPath: generatedAsset.UrlPath,
	//}
	var ExpectedRow []Assets
	ExpectedRow = append(ExpectedRow, Assets{
		ID:        assetID,
		Name:      generatedAsset.Name_2,
		UrlPath:   generatedAsset.UrlPath,
		FileID:    0,
		CreatorID: generatedAsset.CreatorID,
	})
	var pageSize int32 = 5
	var pageNumber int32 = 1
	args := ListAssetByCreatorParams{
		CreatorID: generatedAsset.CreatorID,
		Limit:     5,
		Offset:    (pageNumber - 1) * pageSize,
	}
	assetLists, err := testQueries.ListAssetByCreator(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, assetLists)
	returnedAsset := assetLists[0]
	ExpectedRow[0].FileID = returnedAsset.FileID
	require.Equal(t, ExpectedRow, assetLists)
	cleanup(t, generatedAsset)
}

func TestQueries_DeleteAsset(t *testing.T) {
	generatedAsset, _ := createRandomAsset(t)
	args := DeleteAssetParams{
		UrlPath:   generatedAsset.UrlPath,
		CreatorID: generatedAsset.CreatorID,
	}
	err := testQueries.DeleteAsset(context.Background(), args)
	require.NoError(t, err)
	asset2, err := testQueries.GetAssetDetailsByUrl(context.Background(), generatedAsset.UrlPath)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, asset2)
}
