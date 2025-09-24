# Honeycomb Derived Column Validator

A command-line tool to validate and test Honeycomb derived column expressions.

## Overview

This tool helps validate derived column expressions used in Honeycomb by parsing them and checking for syntax errors. It can be used to test expressions before adding them to your Honeycomb datasets.

## Running it on the command line

### Command Usage

After installing the tool via:

    go install github.com/honeycombio/honeycomb-derived-column-validator

You can run the validator with:

    echo "expression" | honeycomb-derived-column-validator

  or

    honeycomb-derived-column-validator -f expression.txt

### Available Flags

- `-f`        Read from a file instead of stdin
- `-v`        Enable debug mode, which outputs all the detected syntax
- `-help`     Display a help message with usage instructions.

### Example Runs

#### Successful Run

Run the command with a valid expression:

    echo "IF(EQUALS(\$http.response.status_code, 200), 1)" | honeycomb-derived-column-validator -v

Output:

    IF(EQUALS($http.response.status_code, 200), 1)
    operator:  IF
    operator:  EQUALS
    column:  http.response.status_code
    literal:  200
    literal:  1

and the command will exit with a status of 0

#### Failed Run – Mismatched Parenthesis

Run the command with a typoed expression (missing a closing paren):

    echo "IF(1,2" | honeycomb-derived-column-validator -v

Output:

    IF(1,2
    operator:  IF
    Error parsing derived column: mismatched input '<EOF>' expecting {',', ')'}
    exit status 1


## Running it in Terraform

The command automatically detects if has received JSON via stdin, intended for use with Terraform's [external data source](https://registry.terraform.io/providers/hashicorp/external/latest/docs/data-sources/external)

The expected input should look like this:

    {"expression": "IF(EQUALS(\$http.response.status_code, 200), 1)"}

And the expected output will be usable as `data.external.resourcename.result.expression`

See the [Terraform examples](examples/terraform/external_data) for more.

## AI Assistant Support

Want to use AI to help write derived columns? Check out [LLM_PROMPT_USAGE.md](LLM_PROMPT_USAGE.md) for instructions on using the included prompt with Cursor, Claude.ai, Claude Code, ChatGPT, Google Gemini, and other AI assistants.
