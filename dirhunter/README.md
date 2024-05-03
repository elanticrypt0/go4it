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