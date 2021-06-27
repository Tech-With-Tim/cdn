package cache

import (
	"context"
	"database/sql"
	db "github.com/Tech-With-Tim/cdn/db/sqlc"
	"github.com/Tech-With-Tim/cdn/utils"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
)

var postCache PostCache
var testQueries *db.Queries

func createRandomAsset(t *testing.T) (db.CreateAssetParams, int64) {
	arg := db.CreateAssetParams{
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
func cleanup(t *testing.T, asset db.CreateAssetParams) {
	deleteArgs := db.DeleteAssetParams{
		UrlPath:   asset.UrlPath,
		CreatorID: asset.CreatorID,
	}
	err := testQueries.DeleteAsset(context.Background(), deleteArgs)
	require.NoError(t, err)
}

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../", "test")
	if err != nil {
		log.Fatalln(err.Error())
	}
	if config.RedisHost == "" {
		os.Exit(0)
	}
	dbSource := utils.GetDbUri(config)
	testDB, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatalln(err.Error())
	}
	testQueries = db.New(testDB)
	postCache = NewRedisCache(config.RedisHost, config.RedisDb, config.RedisPass, 1)
	os.Exit(m.Run())
}

func TestRedisCache_Set(t *testing.T) {
	var n = 10
	fileRowChan := make(chan db.GetFileRow)
	errChan := make(chan error)
	randAssetChan := make(chan db.CreateAssetParams)
	var fileRow db.GetFileRow
	var err error
	var cachedRow *db.GetFileRow
	var randomAsset db.CreateAssetParams
	for i := 0; i < n; i++ {
		go func() {
			randomAsset, _ := createRandomAsset(t)
			fileRow, err := testQueries.GetFile(context.Background(), randomAsset.UrlPath)
			postCache.Set(randomAsset.UrlPath, &fileRow)
			randAssetChan <- randomAsset
			fileRowChan <- fileRow
			errChan <- err
		}()
	}

	for l := 0; l < n; l++ {
		randomAsset = <-randAssetChan
		fileRow = <-fileRowChan
		err = <-errChan
		require.NoError(t, err)
		require.NotEmpty(t, fileRow)
		require.Equal(t, randomAsset.Data, fileRow.Data)
		cachedRow = postCache.Get(randomAsset.UrlPath)
		require.Equal(t, fileRow, *cachedRow)
		cleanup(t, randomAsset)
	}

}
