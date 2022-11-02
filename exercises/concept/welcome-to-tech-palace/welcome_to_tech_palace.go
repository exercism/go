package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, "+strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var str string
    for i:=0;i<numStarsPerLine;i++{
        str+="*"
        if i==numStarsPerLine-1{
            str+="\n"
        }
    }

    str+=welcomeMsg
    str+="\n"
    for i:=0;i<numStarsPerLine;i++{
        str+="*"
    }
    return str
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    var str string
    var count int
    var end int 
    var start int
	for i,word := range oldMsg{
        if (string(word)>="A" && string(word)<="Z") || (string(word)>="a" && string(word)<="z") || string(word)=="%"{
            if count ==0 {
                start = i
            }
            end = i 
            count++
        }
    }

    for i,word := range oldMsg{
        if i>=start && i<=end{
            str += string(word)
        }
    }
    return str
}
