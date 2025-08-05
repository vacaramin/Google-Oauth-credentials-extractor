package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func main() {
	fmt.Println("🔑 Google Drive OAuth Token Generator")
	fmt.Println("=====================================")

	// Check if credentials file exists
	credentialsPath := "oauth-credentials.json"
	if _, err := os.Stat(credentialsPath); os.IsNotExist(err) {
		fmt.Printf("❌ Error: %s not found\n", credentialsPath)
		fmt.Println("\n📋 Setup Instructions:")
		fmt.Println("1. Go to https://console.cloud.google.com/")
		fmt.Println("2. Navigate to APIs & Services > Credentials")
		fmt.Println("3. Click 'Create Credentials' > 'OAuth 2.0 Client IDs'")
		fmt.Println("4. Choose 'Desktop application'")
		fmt.Println("5. Download the JSON file and save it as 'oauth-credentials.json' in this directory")
		fmt.Println("6. Make sure Google Drive API is enabled in your project")
		return
	}

	// Generate OAuth token
	if err := generateOAuthToken(); err != nil {
		log.Fatalf("❌ Failed to generate OAuth token: %v", err)
	}

	fmt.Println("\n✅ Success! OAuth token generated and saved to 'token.json'")
	fmt.Println("📁 Copy this token.json file to your project's credentials directory")
	fmt.Println("🔄 This token will automatically refresh and never expire (unless manually revoked)")
}

func generateOAuthToken() error {
	// Read OAuth credentials
	credentialsData, err := os.ReadFile("oauth-credentials.json")
	if err != nil {
		return fmt.Errorf("failed to read oauth-credentials.json: %w", err)
	}

	// Parse OAuth config
	config, err := google.ConfigFromJSON(credentialsData, drive.DriveFileScope)
	if err != nil {
		return fmt.Errorf("failed to parse credentials: %w", err)
	}

	// Get token from web
	token, err := getTokenFromWeb(config)
	if err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}

	// Save token to file
	if err := saveToken("token.json", token); err != nil {
		return fmt.Errorf("failed to save token: %w", err)
	}

	// Test the token by making a simple API call
	if err := testToken(config, token); err != nil {
		return fmt.Errorf("token test failed: %w", err)
	}

	return nil
}

func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	// Generate authorization URL
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline, oauth2.ApprovalForce)

	fmt.Printf("\n🌐 Step 1: Open this URL in your browser:\n")
	fmt.Printf("%v\n\n", authURL)

	fmt.Printf("📋 Step 2: After authorizing, copy the authorization code from the browser\n")
	fmt.Printf("🔑 Enter authorization code: ")

	var authCode string
	if _, err := fmt.Scanln(&authCode); err != nil {
		return nil, fmt.Errorf("failed to read authorization code: %w", err)
	}

	fmt.Println("\n⏳ Exchanging authorization code for token...")

	// Exchange authorization code for token
	token, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve token: %w", err)
	}

	return token, nil
}

func saveToken(path string, token *oauth2.Token) error {
	fmt.Printf("💾 Saving token to %s...\n", path)

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("failed to create token file: %w", err)
	}
	defer file.Close()

	// Pretty print JSON for readability
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	return encoder.Encode(token)
}

func testToken(config *oauth2.Config, token *oauth2.Token) error {
	fmt.Println("🧪 Testing token by making a Drive API call...")

	ctx := context.Background()
	client := config.Client(ctx, token)

	// Try to create a Drive service and make a simple call
	service, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return fmt.Errorf("failed to create drive service: %w", err)
	}

	// Test by getting user info about Drive
	about, err := service.About.Get().Fields("user").Do()
	if err != nil {
		return fmt.Errorf("failed to test API call: %w", err)
	}

	fmt.Printf("✅ Token is valid! Authenticated as: %s (%s)\n",
		about.User.DisplayName, about.User.EmailAddress)

	return nil
}
