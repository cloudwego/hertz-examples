# Route
You can learn about how to use hertz route:
* static route
* route group
* use middleware with route group
* parameter route
* use anonymous function or decorator to register routes
* route info

## parameter route:  
Parameters such as ':name' are called `named parameters`, and named parameters only match a single path segment
```shell
Pattern: /hertz/:version

 /hertz/v0                  match
 /hertz/v1                  match
 /hertz/v1/profile          no match
 /hertz/                    no match
```
Parameters such as '*action' are called `wildcard parameters` and they match everything. Therefore, they must be located at the end of the pattern
```shell
Pattern: /src/*filepath

 /src/                     match
 /src/somefile.go          match
 /src/subdir/somefile.go   match
```