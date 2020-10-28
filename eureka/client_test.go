package eureka

import (
	"testing"
)
var inst *Application

func TestClient_Register(t *testing.T) {
	client := NewClient("http://127.0.0.1:8761/eureka")
	inst = client.NewInstance("auth-service", "localhost", "127.0.0.1", 7022, false, "", "", "")
	err := client.Register(inst)
	if err != nil {
		t.Errorf("didn't expect a error but got %v", err.Error())
	}
}

func TestClient_GetAllInstances(t *testing.T) {
	client := NewClient("http://127.0.0.1:8761/eureka")
	instances, err := client.GetAllInstances()
	if err != nil {
		t.Errorf("didn't expect a error but got %v", err.Error())
	}
	if instances == nil {
		t.Errorf("expected instances not to be nil")
	}
	for _, application := range instances.Applications.Application {
		for _,info := range application.InstanceInfo {
			if info.App != "AUTH-SERVICE" {
				t.Errorf("expected app name auth-service but got %v", info.App)
			}
		}
	}

}

func TestClient_GetAllInstancesByApp(t *testing.T) {
	client := NewClient("http://127.0.0.1:8761/eureka")
	instances, err := client.GetAllInstancesByApp(inst)
	if err != nil {
		t.Errorf("didn't expect a error but got %v", err.Error())
	}
	if instances == nil {
		t.Errorf("expected instances not to be nil")
	}
	for _,info := range instances.Instance {
		if info.App != "AUTH-SERVICE" {
			t.Errorf("expected app name auth-service but got %v", info.App)
		}
	}
}

func TestClient_GetInstance(t *testing.T) {
	client := NewClient("http://127.0.0.1:8761/eureka")
	instance, err := client.GetInstance(inst)
	if err != nil {
		t.Errorf("didn't expect a error but got %v", err.Error())
	}
	if instance == nil {
		t.Errorf("expected instances not to be nil")
	}
	if instance.Instance.App != "AUTH-SERVICE" {
		t.Errorf("expected app name auth-service but got %v", instance.Instance.App)
	}
}

func TestClient_DeRegister(t *testing.T) {
	client := NewClient("http://127.0.0.1:8761/eureka")
	err := client.DeRegister(inst)
	if err != nil {
		t.Errorf("didn't expect a error but got %v", err.Error())
	}
}

