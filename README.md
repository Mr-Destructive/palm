# Palm API Go Client

This package provides a Go client for the Palm API. The PaLM API allows developers to build generative AI applications using the PaLM model.

## Installation

Install the package on your system

```
go get github.com/mr-destructive/palm
```

## Usage

To use this package, you'll need a Palm API key. You can get one from the Google Cloud Console.

Set your API key in an .env file:

```
PALM_API_KEY=YOUR_KEY_HERE
```

or in the shell:

```
export PALM_API_KEY=YOUR_KEY_HERE
```

Import the packge with the name `github.com/mr-destructive/palm` as :

```go
package main

import (
  "github.com/mr-destructive/palm"
)

```

### Models

Then you can list models with `ListModels()`:

```go
models, err := palm.ListModels()
if err != nil {
    log.Fatal(err)
}
fmt.Println(models)
```

And get a single model by name with `GetModel(name)`, there are three model names as :

- models/chat-bison-001
- models/text-bison-001
- models/embedding-gecko-001

```go
model, err := palm.GetModel("model/chat-bison-001")
if err != nil {
    log.Fatal(err)
}
fmt.Println(model)
```

### Chat

Get a quick conversation prompt with `ChatPrompt(string)` and further prompts with `Reply(string)` method

```go
chat, err := palm.ChatPrompt("what are you?")
if err != nil {
    panic(err)
}
fmt.Println(chat.Last)
chat.Reply("what can you do!")
fmt.Println(chat.Last)
```

OR fine-tune the `ChatConfig` to the `Chat` method.

```go
chatConfig := palm.ChatConfig{
    Prompt: palm.TextPrompt{
        Text: "what are you?",
    }
    Messages: []palm.Message{
        palm.Message{
            Author: "bot",
            Content: "hello world!",
        },
    },
    Model: "string"
    Context: "string",
    Examples: []palm.Example{
        palm.Example{
            Input: "what are you?",
            Output: "I am a bot",
        }, 
    },
    Temperature: 0.5,
    CandidateCount: 10,
    TopP: 0.5,
    TopK: 10,
}

chat, err := palm.Chat(chatConfig)
if err != nil {
    panic(err)
}
chat.Reply("what can you do!")
```

### Generation

Use the underlying method in the `Chat` methods to get a `ResponseMessage`.

```go
message := palm.MessagePrompt{
    Messages: []palm.Message{
        palm.Message{
            Content: "what is the meaning of life",
        },
    },
}
msgConfig := MessageConfig{
    Prompt: message,
}
m, err := palm.GenerateMessage(msgConfig)
if err != nil {
    fmt.Println(err)
}
fmt.Println(m)

```

## Contributing

Contributions are welcome! Open an issue or submit a PR.

