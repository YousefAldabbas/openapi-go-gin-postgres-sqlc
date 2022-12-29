package structs_test

import (
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/structs"
	"github.com/google/go-cmp/cmp"
)

func TestGetStructKeys(t *testing.T) {
	t.Parallel()

	t.Run("regular keys", func(t *testing.T) {
		t.Parallel()

		ex := Example{}
		want := []string{"key1", "nestedStruct", "nestedStruct.nestedKey", "nestedStruct.nestedStruct2", "nestedStruct.nestedStruct2.nestedKey3"}

		if diff := cmp.Diff(want, structs.GetStructKeys(ex)); diff != "" {
			t.Errorf("MakeGatewayInfo() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("nested keys initialized", func(t *testing.T) {
		t.Parallel()

		ex := Example{}
		// we will want to explicitly initialize what we want, and ignore some fields for project config for instance
		// this way it's clearer when something is missing for some reason since it won't compile.
		ex.NestedStruct.NestedStructArray = append(ex.NestedStruct.NestedStructArray, struct {
			NestedKey           string "json:\"nestedKey\""
			NestedStructInArray struct {
				NestedKey string "json:\"nestedKey\""
			} "json:\"nestedStructInArray\""
		}{})
		want := []string{"key1", "nestedStruct", "nestedStruct.nestedKey", "nestedStruct.nestedStructArray", "nestedStruct.nestedStructArray.nestedKey", "nestedStruct.nestedStructArray.nestedStructInArray", "nestedStruct.nestedStructArray.nestedStructInArray.nestedKey", "nestedStruct.nestedStruct2", "nestedStruct.nestedStruct2.nestedKey3"}

		if diff := cmp.Diff(want, structs.GetStructKeys(ex)); diff != "" {
			t.Errorf("MakeGatewayInfo() mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("pointer field initialized", func(t *testing.T) {
		t.Parallel()

		ex := Example{}
		ex.NestedStructP = &struct {
			NestedKey string "json:\"nestedKey\""
		}{}
		ex.NestedStringArrayP = &[]string{}
		want := []string{"key1", "nestedStruct", "nestedStruct.nestedKey", "nestedStruct.nestedStruct2", "nestedStruct.nestedStruct2.nestedKey3", "nestedStructP", "nestedStructP.nestedKey", "nestedStringArrayP"}

		if diff := cmp.Diff(want, structs.GetStructKeys(ex)); diff != "" {
			t.Errorf("MakeGatewayInfo() mismatch (-want +got):\n%s", diff)
		}
	})
}

type Example struct {
	Key1 string `json:"key1"`

	NestedStruct struct {
		NestedKey         string `json:"nestedKey"`
		NestedStructArray []struct {
			NestedKey           string `json:"nestedKey"`
			NestedStructInArray struct {
				NestedKey string `json:"nestedKey"`
			} `json:"nestedStructInArray"`
		} `json:"nestedStructArray"`
		NestedStruct2 struct {
			NestedKey3 string `json:"nestedKey3"`
		} `json:"nestedStruct2"`
	} `json:"nestedStruct"`

	StringP       *string `json:"stringP"`
	NestedStructP *struct {
		NestedKey string `json:"nestedKey"`
	} `json:"nestedStructP"`
	NestedStringArrayP *[]string `json:"nestedStringArrayP"`
}
