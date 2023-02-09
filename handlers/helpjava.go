package handlers

import "fmt"

func HelpJavaStr(usrs []string) string {

	var fmtStr string
	for _, usr := range usrs {
		fmtStr = fmt.Sprintf("Vish, problemas com Java? Chame esses caras 👇 \n%s\n", usr)
	}

	return fmtStr
}
