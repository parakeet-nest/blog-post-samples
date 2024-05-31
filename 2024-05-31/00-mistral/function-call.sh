#!/bin/bash

: <<'COMMENT'
Chat - Generates a response using the OLLAMA API.

 Args:
    OLLAMA_URL (str): The URL of the OLLAMA API.
    DATA (str): The JSON data to be sent to the API.

 Returns:
    str: The JSON response from the API, containing the generated response and context.
COMMENT
function Chat() {
    OLLAMA_URL="${1}"
    DATA="${2}"

    JSON_RESULT=$(curl --silent ${OLLAMA_URL}/api/chat \
        -H "Content-Type: application/json" \
        -d "${DATA}"
    )
    echo "${JSON_RESULT}"
}


: <<'COMMENT'
Sanitize: Sanitizes the given content by removing any newlines.
It takes one argument, CONTENT, and removes any newline characters (\n) from it using the tr command. 
The sanitized content is then printed to the console.

 Args:
    CONTENT (str): The content to be sanitized.

 Returns:
    str: The sanitized content.
COMMENT
function Sanitize() {
    CONTENT="${1}"
    CONTENT=$(echo ${CONTENT} | tr -d '\n')
    echo "${CONTENT}"
}

: <<'COMMENT'
EscapeDoubleQuotes: Escapes double quotes in the given content by adding a backslash before each double quote.

Args:
  CONTENT (str): The content to escape double quotes in.
    
Returns:
  str: The content with escaped double quotes.
COMMENT
function EscapeDoubleQuotes() {
    CONTENT="${1}"
    CONTENT=$(echo ${CONTENT} | sed 's/"/\\"/g')
    echo "${CONTENT}"
}


OLLAMA_URL=${OLLAMA_URL:-http://localhost:11434}

MODEL="mistral:7b"

read -r -d '' TOOLS_CONTENT <<- EOM
[AVAILABLE_TOOLS]
[
	{
		"type": "function", 
		"function": {
			"name": "hello",
			"description": "Say hello to a given person with his name",
			"parameters": {
				"type": "object", 
				"properties": {
					"name": {
						"type": "string", 
						"description": "The name of the person"
					}
				}, 
				"required": ["name"]
			}
		}
	},
	{
		"type": "function", 
		"function": {
			"name": "addNumbers",
			"description": "Make an addition of the two given numbers",
			"parameters": {
				"type": "object", 
				"properties": {
					"a": {
						"type": "number", 
						"description": "first operand"
					},
					"b": {
						"type": "number",
						"description": "second operand"
					}
				}, 
				"required": ["a", "b"]
			}
		}
	}
]
[/AVAILABLE_TOOLS]
EOM

TOOLS_CONTENT=$(EscapeDoubleQuotes "${TOOLS_CONTENT}")
TOOLS_CONTENT=$(Sanitize "${TOOLS_CONTENT}")


USER_CONTENT='[INST] say "hello" to Bob [/INST]'
USER_CONTENT=$(EscapeDoubleQuotes "${USER_CONTENT}")

read -r -d '' DATA <<- EOM
{
  "model":"${MODEL}",
  "options": {
    "temperature": 0.0,
    "repeat_last_n": 2
  },
  "messages": [
    {"role":"system", "content": "${TOOLS_CONTENT}"},
    {"role":"user", "content": "${USER_CONTENT}"}
  ],
  "stream": false,
  "format": "json",
  "raw": true
}
EOM

jsonResult=$(Chat "${OLLAMA_URL}" "${DATA}")
messageContent=$(echo "${jsonResult}" | jq -r '.message.content')
messageContent=$(Sanitize "${messageContent}")
echo "${messageContent}"


USER_CONTENT='[INST] add 2 and 40 [/INST]'
USER_CONTENT=$(EscapeDoubleQuotes "${USER_CONTENT}")


read -r -d '' DATA <<- EOM
{
  "model":"${MODEL}",
  "options": {
    "temperature": 0.0,
    "repeat_last_n": 2
  },
  "messages": [
    {"role":"system", "content": "${TOOLS_CONTENT}"},
    {"role":"user", "content": "${USER_CONTENT}"}
  ],
  "stream": false,
  "format": "json",
  "raw": true
}
EOM

jsonResult=$(Chat "${OLLAMA_URL}" "${DATA}")
messageContent=$(echo "${jsonResult}" | jq -r '.message.content')
messageContent=$(Sanitize "${messageContent}")
echo "${messageContent}"

#jsonResult=$(Generate "${OLLAMA_URL}" "${DATA}")
#response=$(echo ${jsonResult} | jq -r '.response')
#response=$(Sanitize "${response}")
#echo "${response}"
