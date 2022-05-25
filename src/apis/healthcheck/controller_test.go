package healthcheck

import "testing"

func TestNewController(t *testing.T) {
	want := healthCheckController{healthCheckService{}}
	get := NewController()

	if want != get {
		t.Fail()
	}
}
