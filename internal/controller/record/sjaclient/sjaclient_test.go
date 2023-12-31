package sjaclient

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"

	"git.heb.com/provider-simplejsonapp/apis/records/v1alpha1"
	"github.com/kelseyhightower/envconfig"
)

var env_vars struct {
	Token string
}

var testRecord v1alpha1.RecordParameters = v1alpha1.RecordParameters{
	Name:        "test-record",
	Age:         0,
	Designation: "desg",
	Location:    "loc",
	Todos:       []string{"todo1", "todo2"},
}

func processEnvVars() string {
	envconfig.Process("simplejsonapp", &env_vars)
	return env_vars.Token
}

func createRecord(sjaclient *SjaClient, t *testing.T) v1alpha1.RecordObservation {
	want, err := sjaclient.PostRecord(context.TODO(), testRecord)

	if err != nil {
		t.Errorf("error when creating test record: %s", err.Error())
	}

	return want
}

func deleteRecord(sjaclient *SjaClient, t *testing.T) v1alpha1.RecordObservation {
	want, err := sjaclient.DeleteRecord(context.TODO(), testRecord)

	if err != nil {
		t.Errorf("error when deleting test record: %s", err.Error())
	}

	return want
}

func TestGetRecords(t *testing.T) {
	sjaclient := CreateSjaClient(processEnvVars())

	want := createRecord(sjaclient, t)

	got, err := sjaclient.GetRecords(context.TODO())
	if err != nil {
		t.Errorf("GetRecords() failed with %s", err.Error())
	}

	for _, rp := range got {
		if rp.Name == want.Name {
			if diff := cmp.Diff(rp, want); diff != "" {
				t.Errorf("-got error, +want error, %s", diff)
			}
		}
	}

	deleteRecord(sjaclient, t)
}

func TestGetRecord(t *testing.T) {
	sjaclient := CreateSjaClient(processEnvVars())

	want := createRecord(sjaclient, t)
	got, err := sjaclient.GetRecord(context.TODO(), "test-record")

	if err != nil {
		t.Errorf("GetRecord() failed with %s", err.Error())
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("-got error, +want error, %s", diff)
	}

	deleteRecord(sjaclient, t)
}

func TestPostRecord(t *testing.T) {
	sjaclient := CreateSjaClient(processEnvVars())
	want := createRecord(sjaclient, t)

	var request v1alpha1.RecordParameters
	request.Name = want.Name
	request.Age = want.Age
	request.Designation = want.Designation
	request.Location = want.Location
	request.Todos = want.Todos

	if diff := cmp.Diff(request, testRecord); diff != "" {
		t.Errorf("-got error, +want error, %s", diff)
	}

	deleteRecord(sjaclient, t)
}

func TestPutRecord(t *testing.T) {
	sjaclient := CreateSjaClient(processEnvVars())

	want := createRecord(sjaclient, t)
	want.Designation = "designation changed"

	var request v1alpha1.RecordParameters
	request.Name = want.Name
	request.Age = want.Age
	request.Designation = want.Designation
	request.Location = want.Location
	request.Todos = want.Todos

	got, err := sjaclient.PutRecord(context.TODO(), request)
	if err != nil {
		t.Errorf("PutRecord() failed with %s", err.Error())
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("-want error, +got error: %s", diff)
	}

	deleteRecord(sjaclient, t)
}

func TestDeleteRecord(t *testing.T) {
	sjaclient := CreateSjaClient(processEnvVars())

	want := createRecord(sjaclient, t)

	var request v1alpha1.RecordParameters
	request.Name = want.Name
	request.Age = want.Age
	request.Designation = want.Designation
	request.Location = want.Location
	request.Todos = want.Todos

	got, err := sjaclient.DeleteRecord(context.TODO(), request)
	if err != nil {
		t.Errorf("GetRecord() failed with %s", err.Error())
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("-want error, +got error: %s", diff)
	}

}
