package main

import (
	petv1 "bufdemo/pb/pet/v1"
	"bufdemo/pb/pet/v1/petv1connect"
	"context"
	"net/http"
	"testing"

	"connectrpc.com/connect"
)

func TestClient(t *testing.T) {
	client := petv1connect.NewPetStoreServiceClient(
		http.DefaultClient,
		"http://"+address,
		connect.WithGRPC(),
	)
	req := connect.NewRequest(&petv1.PutPetRequest{
		PetId: "123",
	})
	resp, err := client.PutPet(
		context.Background(),
		req,
	)
	if err != nil {
		t.Fatalf("Failed to call server: %v", err)
	}
	t.Logf("Response: %s", resp.Msg.GetMessage())
	if !resp.Msg.GetSuccess() {
		t.Errorf("Expected success, got %v", resp.Msg.GetSuccess())
	}
}
