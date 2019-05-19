# pushy-penguins

Run with:
```
go run main.go
```
(and will need to add directions for getting dependencies)

Tips:
- `rmdir /s %DIR_NAME%` to remove a directory
- To turn images to byte slices:
    - `go get github.com/hajimehosi/file2byteslice`
    - `file2byteslice -input %INPUT_FILE% -output %OUTPUT_FILE% -package %PACKAGE_NAME% -var %VARIABLE_NAME%`

Dependencies:
- ebiten
- resolv

Notes:
- Oh dear...there's no collision detection...add it to ebiten?
    - They said to try: https://github.com/ByteArena/box2d
        - Seems more robust, but I don't think necessary for Push Penguins...
    - Also found: https://github.com/SolarLune/resolv 
        - I think I need to do this: https://stackoverflow.com/questions/43580131/exec-gcc-executable-file-not-found-in-path-when-trying-go-build