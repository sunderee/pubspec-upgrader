package versions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL string = "https://pub.dev/"

func requestPackageInfo(packageName string) (*packageInfoModel, error) {
	var requestURL string = baseURL + "packages/" + packageName + ".json"
	var httpClient *http.Client = &http.Client{}

	request, requestError := http.NewRequest("GET", requestURL, nil)
	if requestError != nil {
		return nil, &apiException{
			StatusCode: http.StatusBadRequest,
			Message:    requestError.Error(),
		}
	}

	response, responseError := httpClient.Do(request)
	if responseError != nil {
		return nil, &apiException{
			StatusCode: http.StatusInternalServerError,
			Message:    responseError.Error(),
		}
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, &apiException{
			StatusCode: response.StatusCode,
			Message:    response.Status,
		}
	}

	rawBody, parsingError := ioutil.ReadAll(response.Body)
	if parsingError != nil {
		return nil, &apiException{
			StatusCode: http.StatusInternalServerError,
			Message:    parsingError.Error(),
		}
	}
	var packageInfoModel packageInfoModel
	deserializationError := json.Unmarshal(rawBody, &packageInfoModel)
	if deserializationError != nil {
		return nil, &apiException{
			StatusCode: http.StatusInternalServerError,
			Message:    deserializationError.Error(),
		}
	}

	return &packageInfoModel, nil
}

func (exception *apiException) Error() string {
	return fmt.Sprintf("%d: %s", exception.StatusCode, exception.Message)
}
