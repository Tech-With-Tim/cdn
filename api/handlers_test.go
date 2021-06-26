package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	db "github.com/Tech-With-Tim/cdn/db/sqlc"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/Tech-With-Tim/cdn/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"
)

type assetDetailsJsonResponse struct {
	Location string `json:"location"`
	AssetId  string `json:"asset_id"`
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func createAuthToken(exp int64) (string, error) {
	claims := jwt.MapClaims{}

	claims["uid"] = fmt.Sprintf("%v",
		utils.RandomInt(328604827967815690,
			735376244656308274))
	claims["exp"] = exp //time.Now().Add(time.Hour * 24).Unix()
	claims["IssuedAt"] = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	return token.SignedString([]byte(config.SecretKey))
}

func createRandomAsset(t *testing.T, authToken string) (string, *httptest.ResponseRecorder, []byte) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	bytesData := utils.StrToBinary(utils.RandomString(100), 10)
	bytesReader := bytes.NewReader(bytesData)
	formFile, err := writer.CreateFormFile("data", utils.RandomString(4))
	require.NoError(t, err)
	_, err = io.Copy(formFile, bytesReader)
	require.NoError(t, err)
	assetName := utils.RandomString(5)
	_ = writer.WriteField("name", assetName)
	err = writer.Close()
	require.NoError(t, err)
	req, _ := http.NewRequest("POST", "/manage", payload)
	req.Header.Add("Authorization", authToken)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	response := executeRequest(req)
	return assetName, response, bytesData
}

func TestHelloWorld(t *testing.T) {
	req, _ := http.NewRequest("GET", "/testing", nil)

	token, err := createAuthToken(time.Now().Add(time.Hour * 24).Unix())
	require.NoError(t, err)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusUnauthorized, response.Code)
	req.Header.Add("Authorization", token)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body != "Hello World" {
		t.Errorf("Expected Hello World. Got %s", body)
	}
	token, err = createAuthToken(time.Now().Unix() - 60)

	require.NoError(t, err)
	req.Header.Set("Authorization", token)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusUnauthorized, response.Code)
}

func TestCreateAsset(t *testing.T) {
	authToken, err := createAuthToken(time.Now().Add(time.Hour * 24).Unix())
	require.NoError(t, err)
	assetNameChan := make(chan string)
	responseChan := make(chan *httptest.ResponseRecorder)
	fileDataChan := make(chan []byte)
	n := 5
	var res *httptest.ResponseRecorder
	var assetName string

	var fileReq *http.Request
	var fileRes *httptest.ResponseRecorder
	var receivedFileData []byte
	var originalFileData []byte
	var body []byte
	var assetId int
	var AssetDetails db.GetAssetDetailsByIdRow

	for i := 0; i < n; i++ {
		go func() {
			assetN, response, fileData := createRandomAsset(t, authToken)
			assetNameChan <- assetN
			responseChan <- response
			fileDataChan <- fileData
		}()
	}
	for l := 0; l < n; l++ {
		assetName = <-assetNameChan
		res = <-responseChan
		originalFileData = <-fileDataChan

		checkResponseCode(t, http.StatusCreated, res.Code)
		require.NotEmpty(t, res.Body.String())
		body, err = ioutil.ReadAll(res.Body)
		require.NoError(t, err)

		//Check store assets details
		assetResponse := &assetDetailsJsonResponse{}
		err = json.Unmarshal(body, &assetResponse)
		require.NoError(t, err)
		assetId, err = strconv.Atoi(assetResponse.AssetId)
		require.NoError(t, err)
		AssetDetails, err = s.Store.GetAssetDetailsById(context.Background(), int64(assetId))
		require.NoError(t, err)
		require.NotEmpty(t, AssetDetails)
		require.Equal(t, AssetDetails.Name, assetName) //AssetDetails.Name

		//Check stored files byte data
		fileReq, err = http.NewRequest("GET", assetResponse.Location, nil)
		require.NoError(t, err)
		fileRes = executeRequest(fileReq)
		checkResponseCode(t, http.StatusOK, fileRes.Code)
		receivedFileData, err = ioutil.ReadAll(fileRes.Body)
		require.NoError(t, err)
		require.Equal(t, originalFileData, receivedFileData)

		// endpoint /manage/url
		// check if the info is correct
		fileReq, err = http.NewRequest("GET", "/manage/url/"+AssetDetails.UrlPath, nil)
		require.NoError(t, err)
		fileRes = executeRequest(fileReq)
		checkResponseCode(t, http.StatusOK, fileRes.Code)
		receivedFileData, err = ioutil.ReadAll(fileRes.Body)
		require.NoError(t, err)

		manageURLResponse := &db.GetAssetDetailsByUrlRow{}
		err = json.Unmarshal(receivedFileData, manageURLResponse)
		require.NoError(t, err)
		require.Equal(t, assetId, int(manageURLResponse.ID))
		require.Equal(t, assetName, manageURLResponse.Name)

		// not found
		// um, hopefully there isn't a url named * in the test db.
		fileReq, err = http.NewRequest("GET", "/manage/url/*", nil)
		require.NoError(t, err)
		fileRes = executeRequest(fileReq)
		checkResponseCode(t, http.StatusNotFound, fileRes.Code)

		// endpoint /manage/id
		// check if the info is correct
		fileReq, err = http.NewRequest("GET", "/manage/id/"+strconv.Itoa(assetId), nil)
		require.NoError(t, err)
		fileRes = executeRequest(fileReq)
		checkResponseCode(t, http.StatusOK, fileRes.Code)
		receivedFileData, err = ioutil.ReadAll(fileRes.Body)
		require.NoError(t, err)

		manageIDResponse := &db.GetAssetDetailsByIdRow{}
		err = json.Unmarshal(receivedFileData, manageIDResponse)
		require.NoError(t, err)
		require.Equal(t, AssetDetails.UrlPath, manageIDResponse.UrlPath)
		require.Equal(t, assetName, manageIDResponse.Name)
		require.Equal(t, AssetDetails.CreatorID, manageIDResponse.CreatorID)

		// not found
		// um, hopefully there isn't a id named 1 in the test db.
		fileReq, err = http.NewRequest("GET", "/manage/id/1", nil)
		require.NoError(t, err)
		fileRes = executeRequest(fileReq)
		checkResponseCode(t, http.StatusNotFound, fileRes.Code)

		// not int
		fileReq, err = http.NewRequest("GET", "/manage/id/abc", nil)
		require.NoError(t, err)
		fileRes = executeRequest(fileReq)
		checkResponseCode(t, http.StatusBadRequest, fileRes.Code)
	}
}
