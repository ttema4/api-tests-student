package advertisements

import (
	"encoding/json"
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	httpAdvertisements "api-tests-template/internal/client/http/advertisementSearch"
	"api-tests-template/internal/managers/advertisementSearch/models"
)

func GetAdvertisementsList(t *testing.T, token, text string) models.AdvertisementsList {
	raw := httpAdvertisements.HttpGetAdvertisements(t, token, text)
	defer raw.Body.Close()

	body, err := io.ReadAll(raw.Body)
	require.NoError(t, err)

	response := models.AdvertisementsList{}
	require.NoError(t, json.Unmarshal(body, &response))
	return response
}
