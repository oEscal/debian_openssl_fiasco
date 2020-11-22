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
           File where to get the intercepted messages (default "intercepted.txt")
    The default behaviour is to run the decrypt of messages
    ```
