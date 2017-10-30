# utilities
Utilities for Cisco Zeus
# migrate.py
```
git clone https://github.com/CiscoZeus/utilities
```
## Usage
Use python3
```
usage: migrate.py [-h] -s SOURCE_TOKEN -d DEST_TOKEN -o OBJ_TYPE
                  [-r [REPLACE_TOKENS [REPLACE_TOKENS ...]]]

Migrate data from one token to another

optional arguments:
  -h, --help            show this help message and exit
  -s SOURCE_TOKEN, --source_token SOURCE_TOKEN
                        Source token
  -d DEST_TOKEN, --dest_token DEST_TOKEN
                        Dest token
  -o OBJ_TYPE, --obj_type OBJ_TYPE
                        Object type: dashboard OR visualization OR savedsearch
  -r [REPLACE_TOKENS [REPLACE_TOKENS ...]], --replace_tokens [REPLACE_TOKENS [REPLACE_TOKENS ...]]
                        Replace tokens in the format
                        SOURCE_STRING::DEST_STRING

```
The "-r" option is needed because occasionally some strings are different across indices, for example usernames referenced in the tags
## Example
```
python3 ./migrate.py -s <SOURCE_TOKEN> -d <DEST_TOKEN> -r <SOURCE_USERNAME>::<DEST_USERNAME> -o visualization
```
```
python3 ./migrate.py -s <SOURCE_TOKEN> -d <DEST_TOKEN> -r <SOURCE_USERNAME>::<DEST_USERNAME> -o dashboard
```
```
python3 ./migrate.py -s <SOURCE_TOKEN> -d <DEST_TOKEN> -r <SOURCE_USERNAME>::<DEST_USERNAME> -o savedsearch
```
