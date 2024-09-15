package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	llm, err := ollama.New(ollama.WithModel("llama3.1"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	output, err := llm.GenerateContent(ctx, []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, "Ballon d'Or is an annual football award presented by France Football. It has been awarded since 1956, although between 2010 and 2015, an agreement was made with FIFA, and the award was temporarily merged with the FIFA World Player of the Year award and known as the FIFA Ballon d'Or. However, the partnership ended in 2016, and the award was reverted to the Ballon d'Or, while FIFA also reverted to its own separate annual award. The award is presented to the best"),
		llms.TextParts(llms.ChatMessageTypeHuman, "Who is win Ballon d'Or 2021?"),
	})
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(output.Choices); i++ {
		fmt.Println(output.Choices[i].Content)
	}
}
