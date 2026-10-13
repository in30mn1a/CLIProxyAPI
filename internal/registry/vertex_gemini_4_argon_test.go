package registry

import "testing"

func TestGetGeminiVertexModelsIncludesGemini4Argon(t *testing.T) {
	for _, m := range GetGeminiVertexModels() {
		if m != nil && m.ID == "gemini-4-argon" {
			return
		}
	}
	t.Fatal("expected gemini-4-argon in the Vertex static model list")
}
