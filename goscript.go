// Purpose: Compile and run a go file as a script.
//   Put a shebang line like
//     #!/usr/bin/env go  
//   at the top of your one-line program, make it
//   executable, and then you can run it as a script.
//
// Input: Path to a go file, additional arguments.
// Output: Any output from the go script.
// Side effects: Temp files from compiling.
// Author: Issac Trotts <issac.trotts@gmail.com>

package main

import (
    "bufio"
    "exec"
    "flag"
    "fmt"
    "os"
    "path"
    "rand"
)

func printUsage() {
    fmt.Printf("usage: go program.go args...\n")
}

func DropExt(filename string) string {
    ext := path.Ext(filename)
    return filename[:len(filename) - len(ext)]
}

func BaseName(filepath string) string {
    _, filename := path.Split(filepath)
    return DropExt(filename)
}

func IsShebang(line []byte) bool {
    return len(line) >= 2 && line[0] == '#' && line[1] == '1'
}

func StripShebang(input_path string, output_path string) {
    fmt.Printf("Stripping shebang\n")

    infile, err := os.Open(input_path, os.O_RDONLY, 0)
    if err != nil {
        fmt.Printf("Could not open %s for reading.\n", input_path)
        os.Exit(1)
    }
    defer infile.Close() 

    outfile, err := os.Open(output_path, os.O_APPEND | os.O_CREATE, 777)
    if err != nil {
        fmt.Printf("Could not open %s for writing.\n", output_path)
        os.Exit(1)
    }
    defer outfile.Close()

    // Either skip the first line or not depending on whether it is a shebang
    // (#!)
    buf_reader := bufio.NewReader(infile)
    first_line, err := buf_reader.ReadSlice('\n')
    if err != nil {
        fmt.Printf("Failed to read first line.\n")
        os.Exit(1)
    }
    if IsShebang(first_line) {
        // Skip the first line
    } else {
        outfile.Write(first_line)
    }

    for {
        bytes, err := buf_reader.ReadSlice('\n')
        if err == os.EOF {
            break
        }
        outfile.Write(bytes)
    }

    fmt.Printf("Stripping shebang done\n")
}

func main() {
    flag.Parse()
    if flag.NArg() != 1 {
        printUsage()
        os.Exit(1)
    }
    script_path := flag.Arg(0)
    progname := BaseName(script_path)

    workdir := fmt.Sprintf("/tmp/%s-%d", progname, rand.Int())
    cleaned_script_path := fmt.Sprintf("%s/%s.go", workdir, progname)
    // TODO(issac): Handle .8 extensions
    object_file := fmt.Sprintf("%s/%s.6", workdir, progname)

    os.Mkdir(workdir, 0700)
    StripShebang(script_path, cleaned_script_path)

    envv := []string {}
    dir := "."

    argv := []string {"-o", object_file, cleaned_script_path}
    _, _ = exec.Run("6g", argv, envv, dir,
                    exec.PassThrough, exec.PassThrough, exec.PassThrough)

    argv = []string {"-o", progname, object_file}
    _, _ = exec.Run("6l", argv, envv, dir,
                    exec.PassThrough, exec.PassThrough, exec.PassThrough)

    progpath := fmt.Sprintf("./%s", progname)
    fmt.Printf("Running program %s\n", progpath)
    _, _ = exec.Run(progpath, flag.Args()[1:], envv, dir,
                    exec.PassThrough, exec.PassThrough, exec.PassThrough)
}

