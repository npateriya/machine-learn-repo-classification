#pyFNR

Python bindings for [libFNR](https://github.com/cisco/libfnr) from Cisco, Format Preserving Encryption (FPE) library that uses FNR scheme (Flexible Naor-Reingold) [1]

This library currently support Python2.6 to Python3.4 and provides two classes:
* FNR: libFNR wrapper with methods for enciphering/deciphering strings, integers, bytearrays and raw c_char_Arrays. 
* FNR2: FNR wrapper with cycle walking [2] method for extending FNR enciphering scheme to all size of domains < 2^128, not only for sizes which are powers of two (2^block_size).

Library pyFNR also provides one module:
* Util: this module contains classes with various common formats for FPE. Format can be represented as a regular language described by a DFA. For each format this module contains separate class with rank() and unrank() methods for converting words from desired regular language to integers and vice versa. Base class FPE_Format implements rank-then-encipher method from [2]

**IMPORTANT:** This is an experimental module and uses experimental cipher, not for production yet.

###Examples
```
fnr = pyFNR.FNR(key='password', tweak='tweak-is-string', block_size=64)
plain_int = 47
cipher_int = fnr.encrypt_int(plain_int)
plain2_int = fnr.decrypt_int(cipher_int)

plain_str = "Hello"
cipher_str = fnr.encrypt_str(plain_str)
plain2_str = fnr.decrypt_str(cipher_str)

fnr.close()
```

###Installation
Please, install first [libFNR](https://github.com/cisco/libfnr) from Cisco.
Then, install pyFNR as superuser:
```
# python setup.py install
```
or install pyFNR without superuser privileges:
```
$ python setup.py install --user
```

###Run tests, benchmarks and examples
```
$ python tests.py
$ for benchmark in benchmarks/*.py; do python $benchmark; done
$ for example in examples/*.py; do python $example; done
```

*Note:* due to benchmarks on Python3.2 and newer it seems that native `int.to_bytes()` is slower than `bytearray.fromhex()` method, so I choose later method instead of former.

some benchmarks: (average time of 100.000 calls)
with `int.to_bytes()`
```
$ python3.4 benchmarks/benchmark_conversions.py 
_int_to_bytes : 0.004318721294403076ms
_int_to_bytes2: 0.013562514781951904ms
_bytes_to_int : 0.0025879740715026855ms
_bytes_to_int2: 0.010645673274993897ms
encrypt_bytes : 0.02397698163986206ms
decrypt_bytes : 0.02381234645843506ms
```
with `bytearray.fromhex()`
```
$ python3.4 benchmarks/benchmark_conversions.py 
_int_to_bytes : 0.0032341837882995605ms
_int_to_bytes2: 0.01392204999923706ms
_bytes_to_int : 0.0027048635482788087ms
_bytes_to_int2: 0.010532338619232178ms
encrypt_bytes : 0.023919179439544677ms
decrypt_bytes : 0.023606929779052734ms
```

###Bibliography
* [1] Scott Fluhrer Sashank Dara. FNR: Arbitrary length small domain block cipher proposal. Cryptology ePrint Archive, Report 2014/421, 2014. https://eprint.iacr.org/2014/421.pdf
* [2] Mihir Bellare, Thomas Ristenpart, Phillip Rogaway, and Till Stegers. Format-Preserving Encryption. In Selected Areas in Cryptography, pages 295–312, 2009. https://eprint.iacr.org/2009/251.pdf
