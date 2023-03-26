# go-openai-cli
CLI to work with go-openai library

# Getting started 
Build the app

	go build -o /tmp/openai-cli

Running the different commands without any parameters triggers 
examples.

Given that your token is: MY_TOKEN

### Running example ChatGPT command

In a terminal, type: 

    API_TOKEN=MY_TOKEN /tmp/openai-cli ChatGPT

Would print something like: 

> ``` Stream response: dolor sit amet, consectetur adipiscing elit. Nullam tincidunt arcu tellus, vel```

### Running example Dalle-2 command

In a terminal, type:
    
     /tmp/openai-cli dalle2

That will print out a URL where you can view the image. A file called 
example.png will also be generated.