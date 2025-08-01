package techpalace
import  "strings"


// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    return "Welcome to the Tech Palace, " + strings.ToUpper(customer)

}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    stars := strings.Repeat("*", numStarsPerLine)
    result := stars + "\n" + welcomeMsg + "\n"+ stars
    return result
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
result := strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
    return result
}
