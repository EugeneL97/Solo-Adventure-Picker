package services

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/genai"
)

type AdventureAgent struct {
	client *genai.Client
	ctx    context.Context
}

type UserProfile struct {
	Preferences []string `json:"preferences"`
	PastTypes   []string `json:"pastTypes"`
	EnergyLevel string   `json:"energyLevel"` // "low", "medium", "high"
}

// NewAdventureAgent creates a new agent for adventure enhancement
func NewAdventureAgent(ctx context.Context) (*AdventureAgent, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY environment variable not set")
	}

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create Gemini client: %w", err)
	}

	return &AdventureAgent{
		client: client,
		ctx:    ctx,
	}, nil
}

// EnhanceDescription takes a basic adventure and makes it more personalized
func (a *AdventureAgent) EnhanceDescription(adventureName, adventureType, basicDescription string, userProfile *UserProfile) (string, error) {
	prompt := a.buildEnhancementPrompt(adventureName, adventureType, basicDescription, userProfile)

	parts := []*genai.Part{
		{Text: prompt},
	}

	result, err := a.client.Models.GenerateContent(a.ctx, "gemini-2.0-flash", []*genai.Content{
		{Parts: parts},
	}, &genai.GenerateContentConfig{
		Temperature: genai.Ptr(float32(0.7)), // Creative but not too wild
	})

	if err != nil {
		log.Printf("Gemini API error: %v", err)
		return basicDescription, nil // Fallback to original description
	}

	if len(result.Candidates) == 0 || len(result.Candidates[0].Content.Parts) == 0 {
		log.Printf("Empty response from Gemini")
		return basicDescription, nil
	}

	enhancedDescription := result.Candidates[0].Content.Parts[0].Text
	return enhancedDescription, nil
}

// buildEnhancementPrompt creates a personalized prompt for the AI
func (a *AdventureAgent) buildEnhancementPrompt(name, adventureType, description string, profile *UserProfile) string {
	return fmt.Sprintf(`
You are an adventure guide helping personalize activities for a solo explorer. 

Adventure: %s
Type: %s
Basic Description: %s

User Profile:
- Enjoys: %v
- Past adventure types: %v  
- Current energy level: %s

Task: Rewrite the description to be more engaging and personalized for this user. Make it:
1. Exciting and motivational 
2. Specific to their interests
3. Matched to their energy level
4. 2-3 sentences maximum
5. Written in second person ("you will...")

Enhanced Description:`,
		name,
		adventureType,
		description,
		profile.Preferences,
		profile.PastTypes,
		profile.EnergyLevel,
	)
}

// Close cleans up the client connection
func (a *AdventureAgent) Close() error {
	// The genai client doesn't require explicit closing
	return nil
}
