package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

// EnhancedResponse is the structure for the response from OpenAI
type EnhancedResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// GetEnhancedNote enhances the user's note content based on specific rules
func GetEnhancedNote(content string) (string, error) {
	apiURL := "https://api.openai.com/v1/chat/completions"

	// Set up the request payload
	payload := map[string]interface{}{
		"model": "gpt-4", // Model adını doğru girin
		"messages": []map[string]string{
			{
				"role": "system",
				"content": `I’d like you to enhance and refine the notes I provide. My guidelines are as follows, and they must be strictly adhered to:

			1-	The notes I provide may be incomplete or contain just a few keywords. Your task is to interpret what I’m aiming to convey, complete the thoughts, and expand on them.
			2-	The notes can be on any topic.
			3-	I want the notes refined, distilled, and polished. Avoid providing unnecessary information. Keep explanations concise, ideally just a few sentences without excessive detail.
			4-	Ensure that you stay on-topic and avoid deviating from the subject of my notes.
			5-	Write the notes in a 'best practice' format.
			6-	When I enter any notes or keywords, start responding immediately.`,
			},
			{
				"role":    "user",
				"content": content,
			},
		},
	}

	// Encode payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to get a successful response from OpenAI")
	}

	// Decode the response
	var response EnhancedResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	// Ensure we have a response
	if len(response.Choices) == 0 || response.Choices[0].Message.Content == "" {
		return "", errors.New("no suggestions received from OpenAI")
	}

	return response.Choices[0].Message.Content, nil
}