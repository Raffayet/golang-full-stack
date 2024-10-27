package api

import (
	"context"
	"data-merging/internal/service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// SpeechRecognitionHandler represents the handler for speech recognition
type SpeechRecognitionHandler struct {
	speechRecognitionService service.SpeechRecognitionService
}

// NewSpeechRecognitionHandler creates a new handler for speech recognition
func NewSpeechRecognitionHandler(speechRecognitionService service.SpeechRecognitionService) *SpeechRecognitionHandler {
	return &SpeechRecognitionHandler{speechRecognitionService: speechRecognitionService}
}

// UploadAudioHandler handles POST requests for uploading an audio file
func (h *SpeechRecognitionHandler) UploadAudioHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Limit upload size
	r.ParseMultipartForm(10 << 20) // 10 MB limit for uploaded files

	// Retrieve the file from the form data
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to retrieve file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Print some debug info about the file
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)

	// Save the file as output.wav
	tempFile, err := os.Create("output.wav")
	if err != nil {
		http.Error(w, "Failed to create file on server", http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	// Now that the file is saved, process it through the service
	recognizedText, err := h.speechRecognitionService.ProcessAudio(ctx, "output.wav")
	if err != nil {
		http.Error(w, "Unable to process audio", http.StatusInternalServerError)
		return
	}

	// Respond with the recognized text in JSON format
	response := map[string]string{"recognized_text": recognizedText}
	json.NewEncoder(w).Encode(response)
}
