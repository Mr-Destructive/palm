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

```
models/chat-bison-001
models/text-bison-001
models/embedding-gecko-001
```

```go
model, err := palm.GetModel("model/chat-bison-001")
if err != nil {
    log.Fatal(err)
}
fmt.Println(model)
```

### Chat

```go
chatConfig := palm.ChatConfig{
    Messages: []palm.Message{
        palm.Message{
            Content: "what are you!",
        },

    },
}

chat, err := palm.Chat(chatConfig)
if err != nil {
    panic(err)
}
chat.Reply(palm.Message{Content: "what can you do!"})
```

OR a shorter version with `ChatPrompt(string)`

```go
chat, err := palm.ChatPrompt("what are you?")
if err != nil {
    panic(err)
}
fmt.Println(chat.Last)
chat.Reply(palm.Message{Content: "what can you do!"})
fmt.Println(chat.Last)
```


### Generation

```go
message := palm.MessagePrompt{
    Messages: []palm.Message{
        palm.Message{
            Content: "what is the meaning of life",
        },
    },
}
m, err := palm.GenerateMessage(message, map[string]string{"model": "chat-bison-001"})
if err != nil {
    fmt.Println(err)
}
fmt.Println(m)

```

## Contributing

Contributions are welcome! Open an issue or submit a PR.


