package twofer

import "fmt"

// ShareWith function print composible string
func ShareWith(name string) string {
	if len(name) > 0 {
        return fmt.Sprintf("One for %s, one for me.", name)
    }

    return fmt.Sprintf("One for you, one for me.")
}

