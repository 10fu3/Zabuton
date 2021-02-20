package main

import "testing"
import "../Entity"

func TestUserEntity(t *testing.T) {
	a := Entity.
		NewUserBuilder().
		SetUUID("A").
		SetName("B").
		SetThumbnail("C").
		Build()

	if a.GetUUID() != "A"{
		t.Errorf("GetUUID() = %s , want A",a.GetUUID())
	}
	if a.GetName() != "B"{
		t.Errorf("GetName() = %s , want B",a.GetName())
	}
	if a.GetThumbnail() != "C"{
		t.Errorf("GetThumbnail() = %s , want C",a.GetThumbnail())
	}

	b := Entity.EditUser(a)
	c := b.
		 SetUUID("1").
		 SetName("2").
		 SetThumbnail("3").
		 Build()

	if c.GetUUID() != "1"{
		t.Errorf("After edit, GetUUID() = %s , want 1",a.GetUUID())
	}
	if c.GetName() != "2"{
		t.Errorf("After edit, GetName() = %s , want 2",a.GetName())
	}
	if c.GetThumbnail() != "3"{
		t.Errorf("After edit, GetThumbnail() = %s , want 3",a.GetThumbnail())
	}

}
