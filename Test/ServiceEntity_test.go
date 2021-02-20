package main

import (
	"../Entity"
	"testing"
)

func TestServiceEntity(t *testing.T) {
	a := Entity.
		 NewServiceBuilder().
		 SetUUID("B").
		 SetAdminID("C").
		 SetName("A").
		 SetSubDomain("api").
		 Build()

	if a.GetAdminID() != "C" {
		t.Errorf("GetAdminID() = %s , want C", a.GetAdminID())
	}

	if a.GetUUID() != "B" {
		t.Errorf("GetUUID() = %s , want B", a.GetUUID())
	}

	if a.GetSubDomain() != "api" {
		t.Errorf("GetSubDomain() = %s , want api", a.GetSubDomain())
	}

	if a.GetName() != "A" {
		t.Errorf("GetName() = %s , want A", a.GetName())
	}

	b := Entity.EditService(a)

	c := b.
		SetSubDomain("1").
		SetName("2").
		SetAdminID("3").
		SetUUID("4").
		Build()

	if c.GetSubDomain() != "1"{
		t.Errorf("After edit, GetSubDomain() = %s",c.GetSubDomain())
	}

	if c.GetName() != "2"{
		t.Errorf("After edit, GetName() = %s",c.GetName())
	}

	if c.GetAdminID() != "3"{
		t.Errorf("After edit, GetAdminID() = %s",c.GetAdminID())
	}

	if c.GetUUID() != "4"{
		t.Errorf("After edit, GetUUID() = %s",c.GetUUID())
	}

}
