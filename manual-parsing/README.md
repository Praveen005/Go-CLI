To find out what the test coverage is and
visually see which parts of your code are not tested. Run:

```
go test -coverprofile c.out
```


To see which parts of the code are covered, run the following:
> Note: There seems to an error in Windows PowerShell, so run this in Command Prompt or elsewhere.
```
go tool cover -html=c.out
```
