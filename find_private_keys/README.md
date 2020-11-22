## How to run

 - First, you have to compile the code:
    ```bash
    $ go build main.go utils.go
    ```

 - The, you can run with:
    ```bash
    $ ./main <flags>
    $ ./main -h
    Usage of ./main:
      -p string
            File where are saved the public RSA info (default "rsa_public_info.txt")
      -r string
            File where to save the resultant private and public information (default "info.txt")
    ```
