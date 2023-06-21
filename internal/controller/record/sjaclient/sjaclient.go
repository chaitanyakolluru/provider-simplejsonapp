package sjaclient

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/carlmjohnson/requests"
	v1alpha1 "github.com/chaitanyakolluru/provider-simplejsonapp/apis/records/v1alpha1"
)

const SIMPLE_JSON_APP_BASEURL = "http://localhost:8081/json"
const (
	errGetRecords   = "get /records error"
	errGetRecord    = "get /record/{name} error"
	errPutRecord    = "put /record error"
	errPostRecord   = "post /record error"
	errDeleteRecord = "delete /record error"
)

type SjaClient struct{}

func CreateSjaClient() *SjaClient {
	return &SjaClient{}
}

func (s *SjaClient) GetRecords(ctx context.Context) ([]v1alpha1.RecordParameters, error) {
	var response []v1alpha1.RecordParameters

	err := requests.URL(fmt.Sprintf("%s/records", SIMPLE_JSON_APP_BASEURL)).ToJSON(&response).Fetch(ctx)
	if err != nil {
		return []v1alpha1.RecordParameters{}, errors.Wrap(err, errGetRecords)
	}

	return response, nil
}

func (s *SjaClient) GetRecord(ctx context.Context, name string) (v1alpha1.RecordParameters, error) {
	var response v1alpha1.RecordParameters

	err := requests.URL(fmt.Sprintf("%s/records/%s", SIMPLE_JSON_APP_BASEURL, name)).ToJSON(&response).Fetch(ctx)
	if err != nil {
		return v1alpha1.RecordParameters{}, errors.Wrap(err, errGetRecord)
	}

	return response, nil
}

func (s *SjaClient) PostRecord(ctx context.Context, record v1alpha1.RecordParameters) (v1alpha1.RecordParameters, error) {
	var response v1alpha1.RecordParameters

	err := requests.URL(fmt.Sprintf("%s/record", SIMPLE_JSON_APP_BASEURL)).BodyJSON(&record).ToJSON(&response).Fetch(ctx)
	if err != nil {
		return v1alpha1.RecordParameters{}, errors.Wrap(err, errPostRecord)
	}

	return response, nil
}

func (s *SjaClient) PutRecord(ctx context.Context, record v1alpha1.RecordParameters) (v1alpha1.RecordParameters, error) {
	var response v1alpha1.RecordParameters

	err := requests.URL(fmt.Sprintf("%s/record", SIMPLE_JSON_APP_BASEURL)).Put().BodyJSON(&record).ToJSON(&response).Fetch(ctx)
	if err != nil {
		return v1alpha1.RecordParameters{}, errors.Wrap(err, errPutRecord)
	}

	return response, nil
}

func (s *SjaClient) DeleteRecord(ctx context.Context, record v1alpha1.RecordParameters) (v1alpha1.RecordParameters, error) {
	var response v1alpha1.RecordParameters

	err := requests.URL(fmt.Sprintf("%s/record", SIMPLE_JSON_APP_BASEURL)).Delete().BodyJSON(&record).ToJSON(&response).Fetch(ctx)
	if err != nil {
		return v1alpha1.RecordParameters{}, errors.Wrap(err, errDeleteRecord)
	}

	return response, nil
}
