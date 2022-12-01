package general_goutils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/Jeffail/gabs"
)




// ParseIncoming http request body sets the root property as "data"
func ParseIncoming(r *http.Request) (interface{}, error) {
	log.Println("parse incoming data to json for manipulation")
	body, _ := ioutil.ReadAll(r.Body)
	jsonParsed, err := gabs.ParseJSON(body) // parsed incoming data

	// provide a root property
	if !jsonParsed.ExistsP("data") {
		data := map[string]interface{}{
			"data": jsonParsed.Data(),
		}
		jsonParsed, _ = gabs.Consume(data)
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // close for body re-usability

	var i interface{}
	if err != nil {
		msg := "failed to parse incoming data: " + err.Error()
		fmt.Println(msg)
		return nil, err
	}
	i = jsonParsed.Path("data").Data()
	return i, nil
}

func GetIncomingCaseRoot(r *http.Request, shouldLog bool) (*gabs.Container, error) {
	Logger.Info("parse incoming data to json for manipulation")
	body, _ := ioutil.ReadAll(r.Body)
	jsonParsed, err := gabs.ParseJSON(body) // parsed incoming data

	// provide a root property
	if !jsonParsed.ExistsP("data") {
		data := map[string]interface{}{
			"data": jsonParsed.Data(),
		}
		jsonParsed, _ = gabs.Consume(data)
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // close for body re-usability

	if err != nil {
		msg := "failed to parse incoming data: " + err.Error()
		fmt.Println(msg)
		return nil, err
	}

	if shouldLog {
		Logger.Info("incoming: " + string(jsonParsed.Bytes()))
	}

	return jsonParsed, nil
}

func GetFromByteArray(b []byte) (*gabs.Container, error) {

	jsonParsed, err := gabs.ParseJSON(b) // parsed incoming data

	// provide a root property
	if !jsonParsed.ExistsP("data") {
		data := map[string]interface{}{
			"data": jsonParsed.Data(),
		}
		jsonParsed, _ = gabs.Consume(data)
	}

	if err != nil {
		msg := "failed to parse incoming data: " + err.Error()
		fmt.Println(msg)
		return nil, err
	}
	return jsonParsed, nil
}

func GetIncomingCase(r *http.Request) (*gabs.Container, error) {
	body, _ := ioutil.ReadAll(r.Body)
	jsonParsed, err := gabs.ParseJSON(body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	if err != nil {
		msg := "failed to parse incoming case: " + err.Error()
		Logger.Error(msg)
		return nil, err
	}
	return jsonParsed.Path("data"), nil
}

func ParseIncomingIntoGabsRoot(r *http.Request) (*gabs.Container, error) {
	body, _ := ioutil.ReadAll(r.Body)
	jsonParsed, err := gabs.ParseJSON(body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	if err != nil {
		msg := "failed to parse incoming data: " + err.Error()
		Logger.Error(msg)
		return nil, err
	}
	return jsonParsed, nil
}

func ParseDynamic(r *http.Request) (interface{}, error) {
	body, _ := ioutil.ReadAll(r.Body)
	jsonParsed, err := gabs.ParseJSON(body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	var i interface{}
	if err != nil {
		msg := "failed to parse incoming data: " + err.Error()
		Logger.Error(msg)
		return nil, err
	}
	i = jsonParsed.Path("data.dynamicProperty").Data()
	return i, nil
}

// ParseIncomingDataAndItsDynamicProperty returns incoming data, dynamicProperty and a loggable incoming data
func ParseIncomingDataAndItsDynamicProperty(r *http.Request) (interface{}, interface{}, string, error) {
	incoming, err := ParseIncoming(r)

	if err != nil {
		msg := "failed to parse incoming: " + err.Error()
		Logger.Error(msg)
		return nil, nil, "", err
	}

	str := fmt.Sprintf("%v", incoming)
	dynamicProperty, err := ParseDynamic(r)

	if err != nil {
		msg := "failed to parse dynamicProperty: " + err.Error()
		Logger.Error(msg)
		return nil, nil, "", err
	}

	return incoming, dynamicProperty, str, err

}

func GetRootJsonObjectPreCreated(dynamicPropertyByteArray []byte) *gabs.Container {
	json, err := gabs.ParseJSON(dynamicPropertyByteArray)
	if err != nil {
		msg := "failed to parse pre-existing dynamic property: " + err.Error()
		Logger.Error(msg)
		return nil
	}
	return json
}
