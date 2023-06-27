package palm

import (
	"testing"
)

func TestApiKey(t *testing.T) {
	_, err := loadAPIKey(".env")
	if err != nil {
		t.Errorf("loadApiKey failed: %v", err)
	}
}

func TestMakeRequest(t *testing.T) {
	_, err := makeRequest(API_BASE_URL, "GET", nil)
	if err != nil {
		t.Errorf("makeRequest failed: %v", err)
	}
}

func TestListModels(t *testing.T) {
	models, err := ListModels()
	if err != nil {
		t.Errorf("ListModels failed: %v", err)
	}
	if len(models) == 0 {
		t.Error("ListModels returned no models")
	}
}

func TestGetModel(t *testing.T) {
	model, err := GetModel("text-bison-001")
	if err != nil {
		t.Errorf("GetModel failed: %v", err)
	}
	if model.Name != "models/text-bison-001" {
		t.Errorf("GetModel returned incorrect model: %v", model)
	}
}
