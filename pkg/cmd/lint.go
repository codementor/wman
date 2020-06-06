package cmd

import (
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
)

var (
	lintExample = `  # linting is what we do
  wman lint`
)


// newLintCmd returns a new initialized instance of the lint sub command
func newLintCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "lint",
		Short:   "Linting exercise",
		Example: lintExample,
		RunE: lintTest,
	}

	return cmd
}

type T struct {
	field1 string
	field2 string
}

func (t T) String() string  {
	return fmt.Sprint("%s.%s", t.field1, t.field2)
}

func check(str  string) bool {
	if str == "" {
		return true
	} else {
		return false
	}
}

func lintTest(cmd *cobra.Command, args []string) error {

	var x T
	y := T{
		field1: x.field1,
		field2: x.field2,
	}
	fmt.Println(y)
	fmt.Println("value of check: %v", check(""))

	//Reversable was a good exercise right?
	strs := []string{ "kind: Namespace", "kind:Namespace", "kind: Foo", "kind", "kind:  Namespace"}
	newStrs := []string{}

	if strs != nil && len(strs) != 0 {
		fmt.Println("strs is not empty")
	}

	for _, str := range strs {
		newStrs = append(newStrs, str)
	}

	nsRegex := regexp.MustCompile("kind:\\s*Namespace")
	set := make(map[string]bool)

	for _, str := range newStrs {
		if (nsRegex.MatchString(str)) {
			fmt.Println(str)
		}
	}

	for key, _ := range set {
		fmt.Println(key)
	}

	return nil
}