package number

import "strings"

type number string

func (n *number) Translate(s string, response *string) error {
	*reponse = strings.ToUpper(s)
}
