# lusay

A minimal CLI featuring a philosophical, lazy orange cat. Inspired by cowsay, powered by AI and built with Go.

## Why?

I wanted to build something like **cowsay**, but smarter. I had a cat named Lusay, so I used his name for **lusay**.
It turns your message into a speech bubble with custom AI-generated ASCII art, making terminal output a bit more whimsical.

## Getting Started

### Prerequisites
- Go 1.21+
- A Groq API Key

### Setup
1. Clone the repository.
2. Create a `.env` file in the root directory:
   ```env
   GROQ_API_KEY=your_api_key_here
   GROQ_MODEL=llama-3.3-70b-versatile
   ```

### Run
Build and run the application:
```bash
make build
./bin/lusay
```
