DESCRIPTION
    same - find same files in folder(s)

        same [FLAG]... [DIRECTORY]...

    same traverses the file tree, finds regular
    files with the same content (based on file
    size and MD5), and outputs them in table or
    JSON format.

    If no path specified in [DIRECTORY]..., use
    pwd.

INSTALL
    1. If you have golang installed, use 
       `go install github.com/yibangheng/same@latest`.

    2. Or download the binary directly at
       https://github.com/yibangheng/same/releases.

FLAGS
    -h, --help
            Display help.

        --ignore-empty-file
            Do not count empty file.

    -j, --json
            Print results in JSON format.

        --no-trunc
            Do not truncate the MD5s if in table
            format.
            
            No effect in JSON format since the
            MD5s will never be truncated in this
            format. 

    -r, --recursive
            Scan files recursively (including
            files in subdirectories).

    -v, --version
            Display version.
