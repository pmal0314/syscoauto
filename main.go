package main
 
import (
    "flag"
    "strings"
    "fmt"
    "os"
	"github.com/manifoldco/promptui"
	"golang.org/x/term"
)

func main() {
    var nonInteractive bool
    flag.BoolVar(&nonInteractive, "non-interactive", false, "Run in non-interactive mode with defaults")
    flag.Parse()
 
    args := flag.Args()
    // Check if "install" command is provided
    if len(args) > 0 && args[0] == "install" {
        runInstallation(nonInteractive)
    } else {
        fmt.Println("Usage: tgr-cli [--non-interactive] install")
        os.Exit(1)
    }
}
 
func runInstallation(nonInteractive bool) {
    fmt.Println("Starting TGR CLI installation...")
 
    // If running non-interactively, use defaults
    if nonInteractive {
        fmt.Println("Running in non-interactive mode with defaults...")
        fmt.Println("Using default settings: No additional dependencies")
        fmt.Println("Installation completed successfully!")
        return
    }
 
    // Check if we're running in a real TTY or through pexpect/automation
    if !term.IsTerminal(int(os.Stdin.Fd())) && os.Getenv("CLI_AUTOMATION") != "true" {
        fmt.Println("Detected automation environment. Use --non-interactive flag for automated installs.")
        fmt.Println("Example: go run main.go --non-interactive install")
        os.Exit(1)
    }
 
 
    // First prompt: Install other dependencies
    installDepsPrompt := promptui.Prompt{
        Label:     "Do you want to install other dependencies? (y/N)",
		Default:   "y",
        AllowEdit: true,
        Validate: func(input string) error {
            if input == "" {
                return nil // Allow empty for default
            }
            lower := strings.ToLower(input)
            if lower != "y" && lower != "yes" && lower != "n" && lower != "no" {
                return fmt.Errorf("please enter y/yes or n/no")
            }
            return nil
        },
    }
 
    installDeps, err := installDepsPrompt.Run()
    if err != nil {
        if err == promptui.ErrInterrupt {
            fmt.Println("\nInstallation cancelled by user.")
            os.Exit(0)
        }
        fmt.Printf("Prompt failed: %v\n", err)
        fmt.Println("Try running with --non-interactive flag for automated installs.")
        os.Exit(1)
    }
 
    if strings.ToLower(installDeps) == "y" || strings.ToLower(installDeps) == "yes" {
        // Language selection using promptui Select
        languagePrompt := promptui.Select{
            Label: "Which language do you want to use?",
            Items: []string{"TypeScript", "JavaScript"},
        }
 
        selectedIndex, selectedLanguage, err := languagePrompt.Run()
        if err != nil {
            if err == promptui.ErrInterrupt {
                fmt.Println("\nInstallation cancelled by user.")
                os.Exit(0)
            }
            fmt.Printf("Prompt failed: %v\n", err)
            fmt.Println("Try running with --non-interactive flag for automated installs.")
            os.Exit(1)
        }
 
        fmt.Printf("Selected language: %s\n", selectedLanguage)
        _ = selectedIndex // We can use this if needed
 
        // Username prompt
        usernamePrompt := promptui.Prompt{
            Label: "Enter your username",
AllowEdit: true,
            Validate: func(input string) error {
                // Only validate if user has actually typed something
                if len(input) == 0 {
                    return nil // Don't show error for empty input initially
                }
                if len(strings.TrimSpace(input)) < 1 {
                    return fmt.Errorf("username cannot be empty")
                }
                return nil
            },
        }
 
        username, err := usernamePrompt.Run()
        if err != nil {
            if err == promptui.ErrInterrupt {
                fmt.Println("\nInstallation cancelled by user.")
                os.Exit(0)
            }
            fmt.Printf("Prompt failed: %v\n", err)
            fmt.Println("Try running with --non-interactive flag for automated installs.")
            os.Exit(1)
        }
 
 
 
// Admin name prompt
        adminNamePrompt := promptui.Prompt{
            Label: "Enter your admin name",
AllowEdit: true,
            Validate: func(input string) error {
                // Only validate if user has actually typed something
                if len(input) == 0 {
                    return nil // Don't show error for empty input initially
                }
                if len(strings.TrimSpace(input)) < 1 {
                    return fmt.Errorf("username cannot be empty")
                }
                return nil
            },
        }
 
        adminName, err := adminNamePrompt.Run()
        if err != nil {
            if err == promptui.ErrInterrupt {
                fmt.Println("\nInstallation cancelled by user.")
                os.Exit(0)
            }
            fmt.Printf("Prompt failed: %v\n", err)
            fmt.Println("Try running with --non-interactive flag for automated installs.")
            os.Exit(1)
        }
 
        fmt.Printf("Hello, %s!\n", adminName)
fmt.Printf("Hello, %s!\n", username)
fmt.Println("Installation completed successfully!")
 
    } else {
        fmt.Println("Installation cancelled.")
    }
}