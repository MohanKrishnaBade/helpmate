# helpmate
tool that helps to do your tedious job.
```
Usage:
   . logParser [flags]
 
 Flags:
   -f, --filepath string      logger filepath (default "./test.log")
   -h, --help                 help for logParser
   -l, --limit int            * positive limit will read from the beginning of the file and limit the results
                              * 0 will read from the beginning of the file but there is not limit to stdout the results
                              * negative will read from the end of the file  and limit the results
   -q, --querystring string   search term to retrieve the data from a log file
   -s, --show                 stdout the results if the flag is true. (default true)
 
 Global Flags:
       --config string   config file (default is $HOME/...yaml)
```
