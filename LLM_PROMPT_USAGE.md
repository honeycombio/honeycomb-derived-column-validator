# Using the Honeycomb Derived Column LLM Prompt

This guide explains how to use the Honeycomb Derived Column prompt (`examples/LLM_PROMPT.md`) with various AI assistants to help you create derived column expressions from natural language descriptions.

## Overview

The LLM prompt teaches AI assistants about Honeycomb's derived column syntax, available functions, and best practices. Once configured, you can describe what you want in plain English, and the AI will generate the correct derived column expression.

## Quick Start Example

**Input to AI**: "If the response code is 200 and the duration is less than 1000ms, mark it as healthy"

**AI Output**: `IF(AND(EQUALS($response_code, 200), LT($duration_ms, 1000)), "healthy", "unhealthy")`

## Platform-Specific Instructions

### Cursor

1. **Option 1 - Project Context (Recommended)**:
   - Open your project in Cursor
   - Navigate to the `examples/LLM_PROMPT.md` file
   - Press `Cmd+Shift+P` (Mac) or `Ctrl+Shift+P` (Windows/Linux)
   - Select "Cursor: Add File to Chat Context"
   - Now you can ask questions like: "Create a derived column that categorizes response times into fast, medium, and slow buckets"

2. **Option 2 - Direct Reference**:
   - In any chat, reference the file with `@examples/LLM_PROMPT.md`
   - Example: `@examples/LLM_PROMPT.md create a derived column to extract domain from email addresses`

3. **Option 3 - Workspace Instructions**:
   - Create a `.cursorrules` file in your project root
   - Add: `When creating Honeycomb derived columns, refer to examples/LLM_PROMPT.md for syntax and function reference`

4. **Using the Validator (Recommended)**:
   - Cursor can build and run the validator directly
   - Ask Cursor: "Build the honeycomb-derived-column-validator and validate this expression: [your expression]"
   - Cursor will:
     ```bash
     go build -o honeycomb-derived-column-validator
     echo 'YOUR_EXPRESSION' | ./honeycomb-derived-column-validator -v
     ```
   - This provides immediate validation and syntax checking

### Claude.ai (Web/Desktop Interface)

1. **Copy the Prompt**:
   - Open `examples/LLM_PROMPT.md` in your editor
   - Copy the entire contents

2. **Set Up Claude**:
   - Start a new conversation with Claude
   - Paste the prompt as your first message
   - Claude will respond acknowledging it understands the format

3. **Use It**:
   - Ask your questions: "I need a derived column that checks if a user agent contains 'bot' or 'crawler'"
   - Claude will generate: `OR(CONTAINS($user_agent, "bot"), CONTAINS($user_agent, "crawler"))`

4. **Pro Tip**: Save the conversation and reuse it for future derived column needs

### Claude Code

1. **Project-Specific Setup (Recommended)**:
   - Navigate to your project root in terminal
   - Create the commands directory:
     ```bash
     mkdir -p .claude/commands
     ```
   - Copy the prompt to a command file:
     ```bash
     cp examples/LLM_PROMPT.md .claude/commands/honeycomb_derived_columns.md
     ```

2. **Using the Command**:
   - In Claude Code, type `/honeycomb_derived_columns`
   - The full prompt will be loaded into your session
   - Now ask: "Create a derived column to calculate 95th percentile response time"
   - Claude will generate expressions using the loaded syntax

3. **Global Setup (For Multiple Projects)**:
   - Create global commands directory:
     ```bash
     mkdir -p ~/.claude/commands
     cp examples/LLM_PROMPT.md ~/.claude/commands/honeycomb_derived_columns.md
     ```
   - This command will be available in all your Claude Code projects

4. **Direct Reference Alternative**:
   - Without setup, you can reference the file directly:
   - Type: "Read the file at examples/LLM_PROMPT.md"
   - Then: "Using that syntax, create a derived column for [your need]"

5. **Tips**:
   - Command names match the filename (without .md extension)
   - Organize complex prompts in subdirectories: `.claude/commands/honeycomb/`
   - Update the command file when new functions are added

6. **Using the Validator**:
   - Claude Code can build and run the validator
   - Ask: "Build the validator and test this expression: IF(EXISTS($error), 1, 0)"
   - Claude Code will execute:
     ```bash
     go build -o honeycomb-derived-column-validator
     echo 'IF(EXISTS($error), 1, 0)' | ./honeycomb-derived-column-validator -v
     ```
   - You'll get immediate feedback on syntax and see the parsed structure

### ChatGPT

1. **Initial Setup**:
   - Copy the contents of `examples/LLM_PROMPT.md`
   - Start a new chat with ChatGPT
   - Begin with: "I'm going to provide you with documentation about Honeycomb Derived Columns. Please acknowledge when ready."
   - Paste the prompt content
   - ChatGPT will confirm understanding

2. **Creating Custom GPT (ChatGPT Plus)**:
   - Go to "Explore GPTs" → "Create a GPT"
   - Name it "Honeycomb Derived Column Assistant"
   - In Instructions, paste the prompt content
   - Save and use this custom GPT for all derived column needs

3. **Usage**:
   - Ask: "Create a derived column to check if request latency exceeds SLO of 500ms"
   - Response: `GT($request_latency_ms, 500)`

### Google Gemini

1. **Web Interface Setup**:
   - Copy the contents of `examples/LLM_PROMPT.md`
   - Go to [Google AI Studio](https://makersuite.google.com/app/prompts) or [Gemini](https://gemini.google.com)
   - Start a new conversation
   - Paste the prompt with: "Learn this Honeycomb Derived Column syntax reference for our conversation:"
   - Follow with the prompt content

2. **Google AI Studio (Recommended for Developers)**:
   - Create a new "Freeform prompt"
   - In System Instructions, paste the prompt content
   - Save as "Honeycomb Derived Column Generator"
   - Use this saved prompt for consistent results

3. **Usage Examples**:
   - Ask: "Create a derived column that identifies mobile users from user agent"
   - Response: `IF(CONTAINS($user_agent, "Mobile"), "mobile", "desktop")`

4. **Advanced Features**:
   - Use few-shot examples: "Here's what I need: Input: 'error rate over 5%' Output: `GT(DIV($error_count, $total_count), 0.05)`. Now create one for: 'cache hit rate below 80%'"
   - Gemini excels at pattern matching from examples

## Best Practices

### 1. Be Specific About Column Names
- **Good**: "Check if the `http.status_code` column equals 404"
- **Better**: "Check if `$http.status_code` equals 404" (include the $ prefix)
- **Result**: `EQUALS($http.status_code, 404)`

### 2. Describe the Logic Clearly
- **Vague**: "Make a performance indicator"
- **Clear**: "If response time is under 100ms return 'fast', under 500ms return 'normal', otherwise 'slow'"
- **Result**: `IF(LT($response_time_ms, 100), "fast", LT($response_time_ms, 500), "normal", "slow")`

### 3. Mention Data Types When Relevant
- **Example**: "Convert the string field `user_count` to an integer and check if it's greater than 100"
- **Result**: `GT(INT($user_count), 100)`

### 4. Use Examples from Your Data
- **Good**: "Extract the region from a URL like 'https://api-us-east.example.com/v1/users'"
- **Result**: `REG_VALUE($url, \`api-([a-z-]+)\.\`)`

## Common Use Cases

### SLI/SLO Monitoring
```
Request: "Create an SLI that's true when requests are successful (2xx status) and fast (under 300ms)"
Result: AND(GTE($http.status_code, 200), LT($http.status_code, 300), LT($duration_ms, 300))
```

### Data Enrichment
```
Request: "Categorize HTTP status codes into success, client_error, server_error"
Result: IF(LT($http.status_code, 300), "success", LT($http.status_code, 400), "redirect", LT($http.status_code, 500), "client_error", "server_error")
```

### Debugging Aids
```
Request: "Flag requests that took longer than 1 second or have an error field"
Result: OR(GT($duration_ms, 1000), EXISTS($error))
```

### Business Metrics
```
Request: "Calculate revenue per user by multiplying price by quantity, applying 20% discount if premium customer"
Result: IF(EQUALS($customer_type, "premium"), MUL($price, $quantity, 0.8), MUL($price, $quantity))
```

## Troubleshooting

### AI Generates Invalid Syntax
- Remind the AI to use the exact function names from the prompt (all uppercase)
- Ask it to show the expression step-by-step
- For Cursor/Claude Code: Ask them to validate the expression directly
- For other platforms: Copy and validate using the honeycomb-derived-column-validator tool

### AI Confuses Column Names
- Always prefix columns with `$` in your request
- Use quotes for columns with special characters: `$"my column.name"`

### Complex Expressions
- Break down complex logic into steps
- Ask the AI to explain its reasoning
- Request alternative approaches if the first attempt seems complicated

## Validation

After generating expressions, validate them:

```bash
# Install the validator globally
go install github.com/honeycombio/honeycomb-derived-column-validator

# Or build locally if you have the repo
go build -o honeycomb-derived-column-validator

# Test your expression
echo 'IF(EQUALS($status, "success"), 1, 0)' | honeycomb-derived-column-validator -v
```

**Pro Tip for Cursor/Claude Code Users**: These AI assistants can run the validator directly! Just ask them to "validate this expression: [your expression]" and they'll build and run the validator for immediate feedback.

## Advanced Tips

1. **Chain Multiple Requests**: First ask for the expression, then ask for variations or optimizations
2. **Request Explanations**: Ask the AI to explain what each part of the expression does
3. **Get Alternatives**: Ask for multiple ways to achieve the same result
4. **Performance Considerations**: Ask about performance implications for complex expressions

## Support

- **Honeycomb Docs**: [Derived Columns Documentation](https://docs.honeycomb.io/working-with-your-data/derived-columns/)
- **Validator Tool**: [GitHub Repository](https://github.com/honeycombio/honeycomb-derived-column-validator)
- **Community**: Share your experiences and tips in Honeycomb's community forums

Remember: The AI is a tool to help you write expressions faster, but always validate the output and understand what it's doing before using in production!
