# fasta2pansn

This repository was created to contain the code used in [my answer][answer] to the question ["Go file reading, modifing and writting to a new file"][question] asked on Stackoverflow.

> ***Note***
>
> This repo is for demonstration purposes only and will not be supported in any
> way, shape or form. However, you are free to use the code as you wish as per the [license](./LICENSE).

## Build

Simply clone the repo and run `go build` in its root directory.


## Run the tool

```none
Usage of ./fasta2pansn:
  -delimiter string
        Delimiter to use between fields in the pansn output (default "#")
  -fasta string
        Path to the input FASTA file (required)
  -haplotype-id string
        Haplotype ID to use in the pansn output (default "1")
  -sample string
        Sample name to use in the pansn output (default "Sample1")
```

[question]: https://stackoverflow.com/questions/78058903/go-file-reading-modifing-and-writting-to-a-new-file
[answer]: https://stackoverflow.com/a/78059932/1296707
