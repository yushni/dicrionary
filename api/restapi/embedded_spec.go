// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Just a dictionary for \"Галицька\" language.",
    "title": "Dictionary",
    "version": "1.0.0"
  },
  "host": "localhost",
  "basePath": "/",
  "paths": {
    "/words": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "words"
        ],
        "summary": "Get all words",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "default": 20,
            "name": "limit",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "uint64",
            "default": 1,
            "name": "offset",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Word"
              }
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/ErrorBadRequest"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/ErrorgInternalError"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "words"
        ],
        "summary": "Add a new word to the dictionary",
        "operationId": "addWord",
        "parameters": [
          {
            "description": "Metadata of the word that have to be added to the dictionary",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Word"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "required": [
                "id"
              ],
              "properties": {
                "id": {
                  "type": "integer",
                  "format": "uint64"
                }
              }
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/ErrorBadRequest"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/ErrorgInternalError"
            }
          }
        }
      }
    },
    "/words/{wordId}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "words"
        ],
        "summary": "Get word by ID",
        "operationId": "getWord",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "name": "wordId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Word"
            }
          },
          "400": {
            "description": "Invalid Word ID",
            "schema": {
              "$ref": "#/definitions/ErrorBadRequest"
            }
          },
          "404": {
            "description": "Word not found",
            "schema": {
              "$ref": "#/definitions/ErrorNotFound"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/ErrorgInternalError"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "words"
        ],
        "summary": "Delete a word",
        "operationId": "deleteWord",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "name": "wordId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful deletion",
            "schema": {
              "type": "object",
              "required": [
                "id"
              ],
              "properties": {
                "id": {
                  "type": "integer",
                  "format": "uint64"
                }
              }
            }
          },
          "400": {
            "description": "Invalid Word ID",
            "schema": {
              "$ref": "#/definitions/ErrorBadRequest"
            }
          },
          "404": {
            "description": "Word not found",
            "schema": {
              "$ref": "#/definitions/ErrorNotFound"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/ErrorgInternalError"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ErrorBadRequest": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "ErrorNotFound": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "ErrorgInternalError": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "Translation": {
      "type": "object",
      "required": [
        "translation",
        "transcription"
      ],
      "properties": {
        "transcription": {
          "type": "string"
        },
        "translation": {
          "type": "string"
        }
      }
    },
    "Word": {
      "type": "object",
      "required": [
        "word",
        "transcription",
        "translations",
        "origin",
        "samples"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "uint64",
          "readOnly": true
        },
        "origin": {
          "type": "string"
        },
        "samples": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "transcription": {
          "type": "string"
        },
        "translations": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Translation"
          }
        },
        "word": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "name": "words"
    }
  ],
  "externalDocs": {
    "description": "GitHub",
    "url": "https://github.com/yushni/dictionary"
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Just a dictionary for \"Галицька\" language.",
    "title": "Dictionary",
    "version": "1.0.0"
  },
  "host": "localhost",
  "basePath": "/",
  "paths": {
    "/words": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "words"
        ],
        "summary": "Get all words",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "default": 20,
            "name": "limit",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "uint64",
            "default": 1,
            "name": "offset",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Word"
              }
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/ErrorBadRequest"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/ErrorgInternalError"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "words"
        ],
        "summary": "Add a new word to the dictionary",
        "operationId": "addWord",
        "parameters": [
          {
            "description": "Metadata of the word that have to be added to the dictionary",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Word"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "object",
              "required": [
                "id"
              ],
              "properties": {
                "id": {
                  "type": "integer",
                  "format": "uint64"
                }
              }
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/ErrorBadRequest"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/ErrorgInternalError"
            }
          }
        }
      }
    },
    "/words/{wordId}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "words"
        ],
        "summary": "Get word by ID",
        "operationId": "getWord",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "name": "wordId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Word"
            }
          },
          "400": {
            "description": "Invalid Word ID",
            "schema": {
              "$ref": "#/definitions/ErrorBadRequest"
            }
          },
          "404": {
            "description": "Word not found",
            "schema": {
              "$ref": "#/definitions/ErrorNotFound"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/ErrorgInternalError"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "words"
        ],
        "summary": "Delete a word",
        "operationId": "deleteWord",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "name": "wordId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful deletion",
            "schema": {
              "type": "object",
              "required": [
                "id"
              ],
              "properties": {
                "id": {
                  "type": "integer",
                  "format": "uint64"
                }
              }
            }
          },
          "400": {
            "description": "Invalid Word ID",
            "schema": {
              "$ref": "#/definitions/ErrorBadRequest"
            }
          },
          "404": {
            "description": "Word not found",
            "schema": {
              "$ref": "#/definitions/ErrorNotFound"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/ErrorgInternalError"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ErrorBadRequest": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "ErrorNotFound": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "ErrorgInternalError": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "Translation": {
      "type": "object",
      "required": [
        "translation",
        "transcription"
      ],
      "properties": {
        "transcription": {
          "type": "string"
        },
        "translation": {
          "type": "string"
        }
      }
    },
    "Word": {
      "type": "object",
      "required": [
        "word",
        "transcription",
        "translations",
        "origin",
        "samples"
      ],
      "properties": {
        "id": {
          "type": "integer",
          "format": "uint64",
          "readOnly": true
        },
        "origin": {
          "type": "string"
        },
        "samples": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "transcription": {
          "type": "string"
        },
        "translations": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Translation"
          }
        },
        "word": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "name": "words"
    }
  ],
  "externalDocs": {
    "description": "GitHub",
    "url": "https://github.com/yushni/dictionary"
  }
}`))
}
