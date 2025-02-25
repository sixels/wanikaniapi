package wanikaniapi_test

import (
	"net/http"
	"testing"

	"github.com/sixels/wanikaniapi"
	"github.com/sixels/wanikaniapi/wktesting"
	assert "github.com/stretchr/testify/require"
)

func TestStudyMaterialCreate(t *testing.T) {
	client := wktesting.LocalClient()

	_, err := client.StudyMaterialCreate(&wanikaniapi.StudyMaterialCreateParams{
		MeaningNote: wanikaniapi.String("hard"),
		SubjectID:   wanikaniapi.ID(123),
	})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t, `{"study_material":{"meaning_note":"hard","subject_id":123}}`, string(req.Body))
	assert.Equal(t, http.MethodPost, req.Method)
	assert.Equal(t, "/v2/study_materials", req.Path)
	assert.Equal(t, "", req.Query)
}

func TestStudyMaterialList(t *testing.T) {
	client := wktesting.LocalClient()

	_, err := client.StudyMaterialList(&wanikaniapi.StudyMaterialListParams{
		Hidden: wanikaniapi.Bool(true),
		IDs:    []wanikaniapi.WKID{1, 2, 3},
	})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t, "", string(req.Body))
	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "/v2/study_materials", req.Path)
	assert.Equal(t, "hidden=true&ids=1,2,3", wktesting.MustQueryUnescape(req.Query))
}

func TestStudyMaterialGet(t *testing.T) {
	client := wktesting.LocalClient()

	_, err := client.StudyMaterialGet(&wanikaniapi.StudyMaterialGetParams{ID: wanikaniapi.ID(123)})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t, "", string(req.Body))
	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "/v2/study_materials/123", req.Path)
	assert.Equal(t, "", req.Query)
}

func TestStudyMaterialUpdate(t *testing.T) {
	client := wktesting.LocalClient()

	_, err := client.StudyMaterialUpdate(&wanikaniapi.StudyMaterialUpdateParams{
		ID:          wanikaniapi.ID(123),
		MeaningNote: wanikaniapi.String("easy now"),
	})
	assert.NoError(t, err)

	req := client.RecordedRequests[0]
	assert.Equal(t, `{"study_material":{"meaning_note":"easy now"}}`, string(req.Body))
	assert.Equal(t, http.MethodPut, req.Method)
	assert.Equal(t, "/v2/study_materials/123", req.Path)
	assert.Equal(t, "", req.Query)
}
