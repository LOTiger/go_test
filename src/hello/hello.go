package hello

const spanish = "Spanish"
const french = "French"
const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const englishHelloPrefix = "Hello, "

// Hello is a function to say hello
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

    return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
