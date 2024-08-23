package main

import (
	"bytes"
	"fmt"
	"github.com/mailru/easyjson"
	"hw3/model"
	"io"
	"os"
	"strings"
)

func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileContents, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var seenBrowsers []string
	uniqueBrowsers := 0
	foundUsers := ""
	buf := bytes.Buffer{}

	lines := strings.Split(string(fileContents), "\n")

	users := make([]model.User, 0)
	for _, line := range lines {
		user := model.User{}
		// fmt.Printf("%v %v\n", err, line)
		err := easyjson.Unmarshal([]byte(line), &user)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		users = append(users, user)
	}

	for i, user := range users {
		isAndroid := false
		isMSIE := false

		browsers := user.Browsers

		for _, browserRaw := range browsers {
			browser := browserRaw

			if strings.Contains(browser, "Android") {
				isAndroid = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					seenBrowsers = append(seenBrowsers, browser)
					uniqueBrowsers++
				}
			} else if strings.Contains(browser, "MSIE") {
				isMSIE = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					seenBrowsers = append(seenBrowsers, browser)
					uniqueBrowsers++
				}
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		email := strings.ReplaceAll(user.Email, "@", " [at] ")
		buf.WriteString(fmt.Sprintf("[%d] %s <%s>\n", i, user.Name, email))
	}

	foundUsers = buf.String()
	_, err = fmt.Fprintln(out, "found users:\n"+foundUsers)
	if err != nil {
		return
	}
	_, err = fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
	if err != nil {
		return
	}
}

//func FastSearch(out io.Writer) {
//	file, err := os.Open(filePath)
//	if err != nil {
//		panic(err)
//	}
//
//	fileContents, err := io.ReadAll(file)
//	if err != nil {
//		panic(err)
//	}
//
//	var seenBrowsers []string
//	uniqueBrowsers := 0
//	foundUsers := ""
//	buf := bytes.Buffer{}
//
//	lines := strings.Split(string(fileContents), "\n")
//
//	users := make([]map[string]interface{}, 0, len(lines))
//	for _, line := range lines {
//		user := make(map[string]interface{})
//		// fmt.Printf("%v %v\n", err, line)
//		err := json.Unmarshal([]byte(line), &user)
//		if err != nil {
//			if err == io.EOF {
//				break
//			}
//			panic(err)
//		}
//		users = append(users, user)
//	}
//
//	for i, user := range users {
//
//		isAndroid := false
//		isMSIE := false
//
//		browsers, ok := user["browsers"].([]interface{})
//		if !ok {
//			// log.Println("cant cast browsers")
//			continue
//		}
//
//		for _, browserRaw := range browsers {
//			browser, ok := browserRaw.(string)
//			if !ok {
//				// log.Println("cant cast browser to string")
//				continue
//			}
//			if strings.Contains(browser, "Android") {
//				isAndroid = true
//				notSeenBefore := true
//				for _, item := range seenBrowsers {
//					if item == browser {
//						notSeenBefore = false
//					}
//				}
//				if notSeenBefore {
//					seenBrowsers = append(seenBrowsers, browser)
//					uniqueBrowsers++
//				}
//			} else if strings.Contains(browser, "MSIE") {
//				isMSIE = true
//				notSeenBefore := true
//				for _, item := range seenBrowsers {
//					if item == browser {
//						notSeenBefore = false
//					}
//				}
//				if notSeenBefore {
//					seenBrowsers = append(seenBrowsers, browser)
//					uniqueBrowsers++
//				}
//			}
//		}
//
//		if !(isAndroid && isMSIE) {
//			continue
//		}
//
//		email := strings.ReplaceAll(user["email"].(string), "@", " [at] ")
//		buf.WriteString(fmt.Sprintf("[%d] %s <%s>\n", i, user["name"], email))
//	}
//
//	foundUsers = buf.String()
//	_, err = fmt.Fprintln(out, "found users:\n"+foundUsers)
//	if err != nil {
//		return
//	}
//	_, err = fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
//	if err != nil {
//		return
//	}
//}
