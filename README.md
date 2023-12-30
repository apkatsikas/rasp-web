# newhell-web

0. Add `newhell*.json` to the root of the project
0. Add `entities.db` to the root of the project
0. Add the following to `~/.gitconfig`:

    ```
    [url "ssh://git@github.com/"]
        insteadOf = https://github.com/
    ```
0. Add `export GOPRIVATE=github.com/apkatsikas/*` to `~/.bashrc` or equivalent or `go env -w GOPRIVATE=github.com/apkatsikas`
0. Run `ENTITIES_SECRET=mysecret make build-and-run`


## Update
0. Run `go get github.com/apkatsikas/entities@${BRANCH_OR_TAG}` 

## TODOS
* Improve FE, add favico
* Enter for submit on FE - https://stackoverflow.com/questions/20484738/submit-form-on-enter-key-with-javascript
* Graceful shutdown - https://getgophish.com/blog/post/2018-12-02-building-web-servers-in-go/#graceful-shutdown
