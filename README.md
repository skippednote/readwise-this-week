# Readwise This Week

A simple Go CLI tool to list your [Readwise](https://readwise.io/) reading history filtered by tag.

## Features
- Fetches reading history from the Readwise API by tag
- Optionally includes archived items
- Outputs results in a Markdown-friendly format

## Requirements
- Go 1.24+
- A Readwise account and API key

## Setup
1. **Clone the repository:**
   ```sh
   git clone <repo-url>
   cd go-readwise-list
   ```
2. **Install dependencies:**
   ```sh
   go mod tidy
   ```
3. **Create a `.env` file:**
   Add your Readwise API key:
   ```env
   READWISE_API_KEY=your_readwise_api_key_here
   ```

## Usage
Run the CLI with the required `--tag` flag:

```sh
go run cmd/main.go --tag <tag> [--archived=true]
```

- `--tag` (required): The tag to filter reading history by.
- `--archived` (optional): If set to `true`, includes archived items and their locations.

**Example:**
```sh
go run cmd/main.go --tag philosophy --archived=true
```

## Project Structure
- `cmd/main.go`: CLI entry point, parses flags and runs the tool.
- `pkg/readwise/readwise.go`: Readwise API client and core logic.

## Dependencies
- [github.com/joho/godotenv](https://github.com/joho/godotenv): Loads environment variables from `.env`.

## License
MIT
