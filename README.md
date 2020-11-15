# semver

`semver` is a command-line utility to work with the *Semantic Versioning Standards*. This is a wrapper to create the cli tool on top of the semver package by @Masterminds.

## Usage

`semver` is very simple to use command line utility. Simply run the following command to see the usage

```bash
$ semver
``` 

The above will print the usage and available commands as follows.
```shell script
A cli to work with the semver for comparison and incrementing via command line.
For example:

	$ semver inc 1.6.0-alpha.30 alpha

will return 1.6.0-alpha.31

	$ semver inc 1.5.89 patch alpha

will return 1.5.90-alpha

	$ semver greater 1.6.0-alpha.30 1.6.0-alpha

will result in exit(0)

Usage:
  semver [command]

Available Commands:
  equal       Compares the equality of two given semver
  greater     Compare two given semver
  help        Help about any command
  inc         Increments the semver against the given flag
  lesser      Compare two given semver

Flags:
  -h, --help      help for semver
  -v, --verbose   verbose output

Use "semver [command] --help" for more information about a command.
```

### Compare two equal semver strings

```shell script
$ semver equal 1.6.8 1.6.8
```

will result in exit(0) in command line.


```shell script
$ semver equal 4.2.7 1.6.8
```

will result in exit(1) in command line throwing error in the command line.

### Compare greaterThan case

To compare two semver strings LeftHandSide(LHS) and RightHandSide(RHS),

```shell script
$ semver greater 1.4.5 0.3.9
```
Will exit(0) in the command line as safe exit to assure LHS is greater than RHS.
 
 ```shell script
 $ semver greater 0.4.5 1.3.9
 ```
Will exit(1) in the command line which errors in the command-line indicating that RHS is greater than LHS.

### Compare lessThan case

To compare two semver strings LeftHandSide(LHS) and RightHandSide(RHS),

```shell script
$ semver lesser 0.4.5 1.3.9
```
Will exit(0) in the command line as safe exit to assure LHS is lesser than RHS.
 
 ```shell script
 $ semver lesser 1.4.5 1.3.9
 ```
Will exit(1) in the command line which errors in the command-line indicating that RHS is lesser than LHS.

### To increment the semver

To increment the given semver against the given module, i.e., major/minor/patch

```shell script
$ semver inc 1.6.0 major 
```

will result in 
```shell script
2.0.0
```

To create an alpha tag on an incremented version, 

```shell script
$ semver inc 1.6.0 major alpha
```

will result in 

```shell script
2.0.0-alpha
``` 

```shell script
$ semver inc 1.6.0-alpha.8 alpha
```

will result in 

```shell script
1.6.0-alpha.9
``` 

## Thanks
- [Masterminds semver](https://github.com/Masterminds/semver)
- [spf13 cobra](https://github.com/spf13/cobra)
