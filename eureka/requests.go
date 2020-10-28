package eureka

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Request struct {
	Method string
	Path string
	Body []byte
}

func NewRequest(method, path string, body []byte) *Request {
	return &Request{
		Method: method,
		Path:   path,
		Body:   body,
	}
}

func (c *Client) Register(application *Application) error {
	bt, err := json.Marshal(application)
	if err != nil {
		return err
	}
	request := NewRequest("POST", "/apps/" + application.Instance.App, bt)
	response, err := c.SendRequest(request)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusNoContent {
		bd, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("Error: %v, Code: %v", string(bd), response.StatusCode)
	}
	return nil
}

func (c *Client) DeRegister(instance *Application) error {
	request := NewRequest("DELETE", "/apps/" + instance.Instance.App + "/" + instance.Instance.InstanceID, nil)
	response, err := c.SendRequest(request)
	if err != nil {
		return err
	}
	bd, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode == http.StatusOK {
		return nil
	}
	return fmt.Errorf("error: %v, code: %v", string(bd), response.StatusCode)
}

func (c *Client) GetAllInstances() (*Apps, error) {
	request := NewRequest("GET", "/apps",  nil)
	response, err := c.SendRequest(request)
	if err != nil {
		return nil, err
	}
	bd, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %v, code: %v", string(bd), response.StatusCode)
	}
	log.Print(string(bd))
	var applications Apps
	err = json.Unmarshal(bd, &applications)
	if err != nil {
		return nil, err
	}
	return &applications, nil
}

func (c *Client) GetAllInstancesByApp(instance *Application) (*Instances, error) {
	request := NewRequest("GET", "/apps/" + instance.Instance.App,  nil)
	response, err := c.SendRequest(request)
	if err != nil {
		return nil, err
	}
	bd, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %v, code: %v", string(bd), response.StatusCode)
	}
	log.Print(string(bd))
	var instances Instances
	err = json.Unmarshal(bd, &instances)
	if err != nil {
		return nil, err
	}
	return &instances, nil
}

func (c *Client) GetAllInstancesByInstanceID(instance *Application) (interface{}, error) {
	request := NewRequest("GET", "/apps/" + instance.Instance.App + "/" + instance.Instance.InstanceID,  nil)
	response, err := c.SendRequest(request)
	if err != nil {
		return nil, err
	}
	bd, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %v, code: %v", string(bd), response.StatusCode)
	}
	var instances map[string]interface{}
	err = json.Unmarshal(bd, &instances)
	if err != nil {
		return nil, err
	}
	return instances, nil
}

func (c *Client) GetInstance(instance *Application) (*Application, error) {
	request := NewRequest("GET", "/instances/" + instance.Instance.InstanceID,  nil)
	response, err := c.SendRequest(request)
	if err != nil {
		return nil, err
	}
	bd, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %v, code: %v", string(bd), response.StatusCode)
	}
	var instanceInfo Application
	err = json.Unmarshal(bd, &instanceInfo)
	if err != nil {
		return nil, err
	}
	return &instanceInfo, nil
}