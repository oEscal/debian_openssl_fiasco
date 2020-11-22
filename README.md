# debian_openssl_fiasco

## Description (by the Applied Cryptography's teacher)

In the first decade of this century, because of a programming blunder related to the use of pseudo-random numbers, the RSA private data of OpenSSL keys did not have enough variety. In particular, it was possible, and indeed it was relatively common, for RSA moduli that were supposed to be distinct to share a common factor.

The file `rsa_public_info.txt` contains the public data needed to encrypt a message to 20000 distinct individuals. Each of its lines contains three fields:

 - the individual’s name,
 - his/hers RSA 4096-bit modulus, and
 - the corresponding encryption exponent.

Unfortunately, the software that generated the moduli was flawed. Due to a programming bug, it was only able to produce 80000 distinct prime factors (each modulus is a product of two of them).
Your first task is to factor as many moduli as possible (hint: use greatest common divisor computations to find them).In this particular system (supposedly extremely secure due to the 4096-bit moduli), short messages are encoded in a single 4096-bit integer (i.e., in 512 bytes). When Alice wants to send an authenticated short message to Bob she does the following:

 - she composes her plain text numberMin the following way:
    1. the least significant byte of the number is the first character of the message,
    2. the next byte of the number is the second character of the message,
    3. and so on, until the end of the message is reached, which is signaled by putting a byte with a 0 in the number,
    4. the remaining bytes of the number are chosen at random (with some restrictions de-scribed latter on).
   
    For example, to send the message ”Tonight?” the bytes of the plain text number will be
    
    | ... |  rand  |    0   |   '?'  |   't'  |   'h'  |   'g'  |   'i'  |   'n'  |   'o'  |   'T'  |
    |:---:|:------:|:------:|:------:|:------:|:------:|:------:|:------:|:------:|:------:|:------:|
    | ... |  rand  |    0   |   63   |   116  |   104  |   103  |   105  |   110  |   111  |   84   |
    | ... | byte 9 | byte 8 | byte 7 | byte 6 | byte 5 | byte 4 | byte 3 | byte 2 | byte 1 | byte 0 |


    and so the plain text number will be
       M = 84 + 111×256 + 110×256^2 + ··· = 4572394315047661396 +r×256^9
    
    where r is a random number with 512 − 9 = 503 bytes.
   
 - she chooses a random r such that M < N_alice,
 - she then computes M′= M^D_alice mod N_alice
 - if M′ ≥ N_bob she selects another random r and repeats the process
 - she then computes C=M′^E_bob mod N_bob and sendsCto Bob.

The file `intercepted.txt` contains several intercepted messages. Each of its lines contains three fields:
 - the sender’s name,
 - the receiver’s name,
 - the encrypted message.

Your task is to decode as many messages as possible.

## How to run
 - The documentation for each package is inside each of them
