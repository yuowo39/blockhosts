package main

import(
    "fmt"
    "os"
    "bufio"
)

func main() {
    var filePath string
    var maker4 bool = false
    var maker6 bool = false
    var makerh bool = false
    var makerd bool = false

    if len(os.Args) > 1 && len(os.Args) < 7 {
        for i := 1; i < len(os.Args); i++ {
            s := os.Args[i]
            switch s {
            case "-f":
                maker4 = true
                break
            case "-s":
                maker6 = true
                break
            case "-h":
                makerh = true
                break
            case "-d":
                makerd = true
                break
            default:
                filePath = s
                break
            }
        }
        if filePath == "" {
            fmt.Println("The file path is null, exited.")
            return
        }

        if !maker4 && !maker6 && !makerh && !makerd {
            maker4 = true
            makerh = true
        }
        if maker4 && !maker6 && !makerh && !makerd {
            makerh = true
        }
        if !maker4 && maker6 && !makerh && !makerd {
            makerh = true
        }
        if !maker4 && !maker6 && makerh && !makerd {
            maker4 = true
        }
        if !maker4 && !maker6 && !makerh && makerd {
            maker4 = true
        }
        if maker4 && maker6 && !makerh && !makerd {
            makerh = true
        }
        if !maker4 && !maker6 && makerh && makerd {
            maker4 = true
        }

        file, oerr := os.Open(filePath)
        if oerr != nil {
            fmt.Println("Open file error, exited.")
            return
        }
        defer file.Close()

        var lines []string
        var lineCount int = 0
        scanner := bufio.NewScanner(file)
        if scanner.Err() != nil {
            fmt.Println("Read file error, exited.")
            return
        }
        for scanner.Scan() {
            lines = append(lines, scanner.Text())
            lineCount++
        }

        const ipv4h string = "127.0.0.1"
        const ipv6h string = "::1"

        const addtext string = "address=/"
        const ipv4d string = "/127.0.0.1"
        const ipv6d string = "/::1"

        var outputFilePath = filePath + ".output"
        var outputLines []string
        var outputLineCount int = 0

        if maker4 {
            if makerh {
                for i := 0; i < lineCount; i++ {
                    ol := ipv4h + " " + lines[i]
                    outputLines = append(outputLines, ol)
                    outputLineCount++
                }
            }
            if makerd {
                for i := 0; i < lineCount; i++ {
                    ol := addtext + lines[i] + ipv4d
                    outputLines = append(outputLines, ol)
                    outputLineCount++
                }
            }
        }

        if maker6 {
            if makerh {
                for i := 0; i < lineCount; i++ {
                    ol := ipv6h + " " + lines[i]
                    outputLines = append(outputLines, ol)
                    outputLineCount++
                }
            }
            if makerd {
                for i := 0; i < lineCount; i++ {
                    ol := addtext + lines[i] + ipv6d
                    outputLines = append(outputLines, ol)
                    outputLineCount++
                }
            }
        }

        outputFile, cerr := os.Create(outputFilePath)
        if cerr != nil {
            fmt.Println("Create file error, exited.")
            return
        }
        defer outputFile.Close()

        writer := bufio.NewWriter(outputFile)
        for i := 0; i < outputLineCount; i++ {
            writer.WriteString(outputLines[i] + "\n")
        }
        writer.Flush()
        
    } else {
        fmt.Println("The argument is null, exited.")
        fmt.Println("The arguments are:")
        fmt.Println("  -f        Default: IPv4")
        fmt.Println("  -s        IPv6")
        fmt.Println("  -h        Default: Hosts(Windows, macOS, Linux)")
        fmt.Println("  -d        Dnsmasq")
        fmt.Println("  FilePath")
    }
}
