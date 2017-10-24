## Description

An implementation of External Merge Sort in Go.

## Usage

Generating sample files:

```
go run main.go gen --num_lines 123456
```

A file containing `123456` lines of random 10-char strings will be generated. Default file name is `example.txt`, but you can specify target name with `--output` flag.

Sorting a file:

```
go run main.go sort
```

The sorted file will be stored as `sorted.<input_file_name>`, with `example.txt` as defaut input file name. You can specify input file name with `--input`.

You can specify buffer size with `--buffer size` (integer), which specifies the number of lines from input file that go into a single buffer. 

## Limitations

* The code is not properly optimized.
* You can not specify buffer size in bytes.
* There's no tests yet.
* In case of any errors we simply panic.