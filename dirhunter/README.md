# DirHunter

Scann al content in a path

# Setup

```go
    // create a new instace and fetch all the content of the directory
    dhunter:= dirhunter.New("/var/www/files")
    dhunter.Run("")
```

You can refetch or fetch another dir like this:

```go
    // create a new instace and fetch all the content of the directory
    dhunter.Run("/var/www/files2")
```


## Work with the data

```go
	for _, dir := range dhunter.Directories {
		fmt.Printf("%s: \n", dir.Name)
		for _, dirFile := range dir.Files {
			fmt.Printf("    + %s\n", dirFile.Name)
		}
	}
```

Also you can use this function to print the data in console

```go
    dhunter.PrintDirData()
```