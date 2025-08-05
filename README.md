# Google Drive OAuth Token Generator

## Description
This Go command-line tool generates and saves an OAuth 2.0 token for accessing the Google Drive API. It guides users through the OAuth authorization process, exchanges the authorization code for a refreshable token, tests the token with a simple Drive API call, and saves it to a token.json file. The generated token can be used in applications requiring Google Drive API access with the https://www.googleapis.com/auth/drive.file scope.

## Features
- Generates OAuth 2.0 tokens for Google Drive API access.
- Provides step-by-step instructions for obtaining an authorization code.
- Tests token validity by retrieving user information via the Drive API.
- Saves the token to a token.json file with proper formatting.
- Supports automatic token refresh (non-expiring tokens unless revoked).

## Prerequisites
- Go: Version 1.16 or later. [Download Go](https://golang.org/dl/)
- Google Cloud Project: A project with the Google Drive API enabled and OAuth 2.0 credentials configured.
- Credentials File: A oauth-credentials.json file from the Google Cloud Console.

## Installation
1. Clone the repository:
   git clone https://github.com/your-username/google-drive-oauth-token-generator.git
2. Navigate to the project directory:
   cd google-drive-oauth-token-generator
3. Install dependencies:
   go mod tidy
4. Set up OAuth credentials:
   - Visit the [Google Cloud Console](https://console.cloud.google.com/).
   - Go to APIs & Services > Credentials.
   - Click Create Credentials > OAuth 2.0 Client IDs.
   - Select Desktop application.
   - Download the JSON file and save it as oauth-credentials.json in the project directory.
   - Ensure the Google Drive API is enabled in APIs & Services > Library.

## Usage
1. Run the tool:
```bash
   go run .
```

3. Follow the prompts:
   - If oauth-credentials.json is missing, the tool provides setup instructions and exits.
   - Copy the provided URL, open it in a browser, and authorize the application.
   - Paste the authorization code into the terminal.
4. The tool will:
   - Exchange the code for an OAuth token.
   - Test the token with a Drive API call to fetch user info.
   - Save the token to token.json.
5. Copy token.json to your application's credentials directory for use with the Google Drive API.

## Example Output
🔑 Google Drive OAuth Token Generator
=====================================
```bash
🌐 Step 1: Open this URL in your browser:
https://accounts.google.com/o/oauth2/auth?...
📋 Step 2: After authorizing, copy the authorization code from the browser
🔑 Enter authorization code: [paste code here]
⏳ Exchanging authorization code for token...
💾 Saving token to token.json...
🧪 Testing token by making a Drive API call...
✅ Token is valid! Authenticated as: John Doe (john.doe@example.com)
✅ Success! OAuth token generated and saved to 'token.json'
```
## Configuration
- Credentials File: Place oauth-credentials.json in the project root.
- Token File: The tool outputs token.json in the project root, containing access and refresh tokens.

## Notes
- The token uses the https://www.googleapis.com/auth/drive.file scope, limiting access to files created or opened by the app.
- The token includes a refresh token for automatic renewal unless revoked in the Google Cloud Console.
- Ensure your Google Cloud project has the Drive API enabled and a configured OAuth consent screen.

## Contributing
Contributions are welcome! To contribute:
1. Fork the repository.
2. Create a branch:

```
   git checkout -b feature-branch-name
```
4. Commit changes:
```
   git commit -m "Add feature"
```
6. Push the branch:
```
   git push origin feature-branch-name
```
8. Open a pull request.

Follow the [Code of Conduct](CODE_OF_CONDUCT.md) and [Contributing Guidelines](CONTRIBUTING.md).

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact
For questions or feedback:
- Open an issue on the [GitHub repository](https://github.com/vacaramin/Google-Oauth-credentials-extractor/issues).
