package main

import (
	petv1 "bufdemo/pb/pet/v1"
	"bufdemo/pb/pet/v1/petv1connect"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/proto"
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

func BenchmarkMarshal(b *testing.B) {
	b.Run("vtproto", func(b *testing.B) {
		pet := &petv1.PutPetRequest{
			PetId: "123",
		}
		b.ResetTimer()
		for b.Loop() {
			_, err := pet.MarshalVT()
			if err != nil {
				b.Fatalf("Failed to marshal: %v", err)
			}
		}
	})
	b.Run("proto", func(b *testing.B) {
		pet := &petv1.PutPetRequest{
			PetId: "123",
		}
		b.ResetTimer()
		for b.Loop() {
			_, err := proto.Marshal(pet)
			if err != nil {
				b.Fatalf("Failed to marshal: %v", err)
			}
		}
	})
	b.Run("json", func(b *testing.B) {
		pet := &petv1.PutPetRequest{
			PetId: "123",
		}
		b.ResetTimer()
		for b.Loop() {
			_, err := json.Marshal(pet)
			if err != nil {
				b.Fatalf("Failed to marshal: %v", err)
			}
		}
	})

}
