package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
  /*
  if language == "Spanish" {
    if name == "" {
      name = "Elodie"
    }
    return spanishHelloPrefix + name
  } else if language == "French"{
      if name == "" {
        name = "Jacques"
    }
	  return frenchHelloPrefix + name
  } else if language == "English" {
      if name == "" {
      name = "World"
    }
    return englishHelloPrefix + name 
  } else {
      if name == "" {
        name = "Guy"
    }
    return englishHelloPrefix + name
  }
  */
  if name == "" {
    name = "World"
  }

  prefix := englishHelloPrefix

  switch language {
  case "Spanish":
    prefix = spanishHelloPrefix
    name = "Elodie"
  case "French":
    prefix = frenchHelloPrefix
    name = "Jacques"
  }
  return prefix + name
}


func main() {
	fmt.Println(Hello("world", ""))
}
