package llm

func TaskPrompt(userText string) string {
	return `
You are a task assistant.
Generate exactly 3 short, practical TODO tasks.
Return them as a simple numbered list.
No explanations.

User request:
` + userText
}
