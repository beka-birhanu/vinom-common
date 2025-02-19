package gamepb_test

import (
	"testing"

	gamepb "github.com/beka-birhanu/vinom-game-encoder"
	"github.com/google/uuid"
)

func TestGamePb(t *testing.T) {
	t.Run("GamePb_Marshal_Unmarshal_Action", testMarshalUnmarshalAction)
}

func testMarshalUnmarshalAction(t *testing.T) {
	id := uuid.New()
	encoder := &gamepb.Protobuf{}
	action := encoder.NewAction()
	action.SetDirection("North")
	action.SetID(id)

	if action.GetID() == uuid.Nil {
		t.Errorf("Expected ID: %s, got %s", id, action.GetID())
	}
	marshaledAction, err := encoder.MarshalAction(action)
	if err != nil {
		t.Errorf("Expected marshal, got error: %s", err)
		t.FailNow()
	}

	unmarshaledAction, err := encoder.UnmarshalAction(marshaledAction)
	if err != nil {
		t.Errorf("Expected unmarshal, got error: %s", err)
		t.FailNow()
	}

	if unmarshaledAction.GetID() != action.GetID() {
		t.Errorf("Expected ID: %s , got : %s", action.GetID(), unmarshaledAction.GetID())
		t.FailNow()
	}
}
