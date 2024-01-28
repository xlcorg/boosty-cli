package util

import (
	"fmt"
	"os"
	"strings"
)

func VerifyName(name string) error {
	blogName := strings.TrimSpace(name)
	if len(blogName) == 0 {
		return fmt.Errorf("name must be specified")
	}

	return nil
}

func CheckError(err error) {
	if err == nil {
		return
	}

	msg := fmt.Sprintf("Error: %s", err.Error())
	fatal(msg, 1)
}

func fatal(msg string, code int) {
	if len(msg) > 0 {
		// add newline if needed
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}

		_, _ = fmt.Fprint(os.Stderr, msg)
	}
	os.Exit(code)
}

//
//func verify(cmd *cobra.Command, args []string) error {
//	if len(args) != 1 {
//		return fmt.Errorf("invalid arguments")
//	}
//	err := verifyName(args[0])
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func verifyName(name string) error {
//	blogName := strings.TrimSpace(name)
//	if len(blogName) == 0 {
//		return fmt.Errorf("name must be specified")
//	}
//
//	return nil
//}
//
//func checkError(err error) {
//	if err == nil {
//		return
//	}
//
//	msg := fmt.Sprintf("Error: %s", err.Error())
//	fatal(msg, 1)
//}
//
//func fatal(msg string, code int) {
//	if len(msg) > 0 {
//		// add newline if needed
//		if !strings.HasSuffix(msg, "\n") {
//			msg += "\n"
//		}
//
//		fmt.Fprint(os.Stderr, msg)
//	}
//	os.Exit(code)
//}
