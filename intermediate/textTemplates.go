package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {

	tmpl, err := template.New("example").Parse("Welcome , {{.name}}!\nHow are you doing? \n")

	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"name": "John",
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	tmpl2 := template.Must(template.New("example1").Parse("Welcome , {{.name}}!\nHow are you doing? \n"))
	err = tmpl2.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	fmt.Sprintf("----------------------------------------------------------------------------------")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')

	name = strings.TrimSpace(name)

	// Define name templates for different types of
	templates := map[string]string{
		"welcome":      "Welcome, {{.name}}!\n Happy to welcome you!!",
		"notification": "{{.name}}, you have a new notification!! : {{.notification}} \n",
		"error":        "Oops, something went wrong! : {{.errorMessage}} \n",
	}

	// Parse and store templates
	paresedTemplates := make(map[string]*template.Template)
	for name, tmpl := range templates {
		paresedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		fmt.Println("Show Menu")
		fmt.Println("1. Join")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get Error")
		fmt.Println("4. Quit")
		fmt.Println("Choose an Option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]interface{}
		var templ *template.Template

		switch choice {
		case "1":
			templ = paresedTemplates["welcome"]
			data = map[string]interface{}{
				"name": name,
			}
		case "2":
			fmt.Println("Enter your notification message : ")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			templ = paresedTemplates["notification"]
			data = map[string]interface{}{
				"name":         name,
				"notification": notification,
			}
		case "3":
			fmt.Println("Enter your Error message : ")
			errorMessage, _ := reader.ReadString('\n')
			errorMessage = strings.TrimSpace(errorMessage)
			templ = paresedTemplates["error"]
			data = map[string]interface{}{
				"errorMessage": errorMessage,
			}
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
			continue
		}

		err = templ.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println("Error executing template: ", err)
		}

	}

}
