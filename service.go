package grpcsay

import (
	"fmt"
	"strings"
)

const cow = `
 _%v_
< %v >
 -%v-
     \   ^__^
      \  (oo)\_______
         (__)\       )\/\
             ||----w |
             ||     ||
`

const gopher = `
 _%v_
< %v >
 -%v-
   \
    \
  \ʕ◔ϖ◔ʔ/
`

func Say(message string) string {
	top := strings.Repeat("_", len(message))
	bottom := strings.Repeat("-", len(message))

	template := gopher

	return fmt.Sprintf(template, top, message, bottom)
}
