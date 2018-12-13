package api

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/gorilla/mux"
)


// TODO this test isn't really that good but it's here for the sake of testing
func routerTestParams(t *testing.T) {

	const expectedObtainedString = "\n...expected = %#v\n...obtained = %#v"

	var router *requestRouter
	router = &requestRouter{
		mux.NewRouter(),
	}

	var paramArray []string
	routerHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		variables := mux.Vars(r)
		variableOne := variables["var1"]
		variableTwo := variables["var2"]

		paramArray = append(paramArray, variableOne)
		paramArray = append(paramArray, variableTwo)
	})

	router.HandleFunc("/pathOne/{var1}/pathTwo/{var2}", routerHandler).Methods("GET")

	rec := httptest.NewRecorder()

	testPathOne := "pathOne/variableOne/pathTwo/variableTwo"
	req, err := http.NewRequest("GET", testPathOne, nil)
	if err != nil {
		t.Error(err.Error())
	}

	routerHandler.ServeHTTP(rec, req)

	if http.StatusOK != rec.Code {
		t.Errorf(expectedObtainedString, http.StatusOK, rec.Code)
	}

	if paramArray[0] != "variableOne" {
		t.Errorf(expectedObtainedString, "variableOne", paramArray[0])
	}

	if paramArray[1] != "variableOne" {
		t.Errorf(expectedObtainedString, "variableOne", paramArray[1])
	}
}
