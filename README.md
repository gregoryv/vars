Package [vars](https://pkg.go.dev/github.com/gregoryv/vars) provides
copy functions for variables.

## Quick start

    go get github.com/gregoryv/vars
	
	
## Usage

	import github.com/gregoryv/vars
	
	var (
        i int
        s string
    )
	
    _ = vars.Copy(
	    &i, 0, 
        &s, "hello",
	)
	
