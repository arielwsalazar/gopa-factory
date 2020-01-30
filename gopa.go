package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    var m map[string]string
    m = make(map[string]string)

    reader := bufio.NewReader(os.Stdin)
    argv := os.Args[1:]
    argc := len(argv)

    if argc == 0 {
        currentDirectory, _ := os.Getwd()
        fmt.Printf("Dear friend, you are to create the scafolding on default directory %s\n", currentDirectory)
        fmt.Print("are you ok with that? (press 'Y' to continue)-> ")
        text, _ := reader.ReadString('\n')
        text = strings.Replace(text, "\n", "", -1)
        if strings.Compare("Y", text) != 0 {
            fmt.Println("bye bye, see you later.")
            return
        }
        m["--d"] = currentDirectory
    }

    loadParameters(m)

    if !isValidParameters(m) {
        fmt.Println("bye bye, see you later, parameters invalid.")
        return
    }

    directoriesToCreate := []string{
        "cmd",
        "internal",
        "internal/platform",
        "vendor",
    }

    wd, _ := m["--d"]
    fmt.Printf("root directory %s\n", wd)

    for index := 0; index < len(directoriesToCreate); index++ {
        if _, err := os.Stat(wd + "/" + directoriesToCreate[index]); err != nil {
            if os.IsNotExist(err) {
                err2 := os.MkdirAll(wd+"/"+directoriesToCreate[index], 777)
                if err2 != nil {
                    fmt.Printf("there was a problem creating %s \n", wd+"/"+directoriesToCreate[index])
                }
                fmt.Printf("directory %s\n", wd+"/"+directoriesToCreate[index])
            }
        }
    }

}

func isValidParameters(m map[string]string) bool {
    _, ok := m["--d"]
    return ok
}

func loadParameters(m map[string]string) {
    params := os.Args[1:]
    paramSize := len(params)
    for index := 0; index < paramSize; index++ {
        if strings.HasPrefix(params[index], "--") {
            if index+1 < paramSize {
                if !strings.HasPrefix(params[index+1], "--") {
                    m[params[index]] = params[index+1]
                } else {
                    m[params[index]] = ""
                }
            }
        }
    }
}