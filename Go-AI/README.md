## Go AI

GO AI is AI powered CLI build with Cobra and Golang. It answers your question from the terminal. It uses Google Gemini API to get the answer of your question.


### Installation

First we need to get the API key from the Google Gemini API. To get it, head over here [https://aistudio.google.com/app/apikey](https://aistudio.google.com/app/apikey) and get that. It's FREE and you get that in 30 Seconds. Once you get the API key set is an environment variable by executing the following command:

```sh
export GEMINI_API_KEY=<YOUR_API_KEY>
```

Note: The issue with this method is that the environment variable will only exist for the current session as you close the terminal it's gone. To avoid this issue add the **export** command to a shell configuration file, such as `.bashrc`, `.bash_profile`, or `.zshrc` (depending on your shell). In this way, you can access the CLI from anywhere in the system.

Once you have the API key set, you can use the cli by running the following command:

```sh
go run main.go
```

To use the search feature, you can use the following command:

```sh
go run main.go search "YOUR QUESTION"
```
