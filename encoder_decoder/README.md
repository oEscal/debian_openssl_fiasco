## How to run

 - First, you have to compile the code:
    ```bash
    $ go build main.go decrypt.go encrypt.go utils.go
    ```

 - The, you can run with:
    ```bash
    $ ./main <flags>
    $ ./main -h
    Usage of ./main:
      -e    Encrypt message
      -i string
            File where to get the public and private info of each user (default "info.txt")
      -m string
            File where to get the intercepted messages (just valid for decrypt) (default "intercepted.txt")
      -message string
            Message to encrypt (just valid for encrypt)
      -receiver string
            Message receiver (just valid for encrypt) (default "Charlie Brown")
      -sender string
            Message sender (just valid for encrypt) (default "Brian York")
    The default behaviour is to run the decrypt of messages
    ```
