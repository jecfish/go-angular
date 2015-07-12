#Demo App for Angular 1.4 + GO

Get, Add, Delete friends

Install: Enviroment
=======
1. Install go: https://golang.org/
2. Install node js: https://nodejs.org/
3. Install environment variables:-
  * For Windows user: create a GOPATH system variable by go to Control Panel > System > Advance system settings > Advanced tab > Environment Variables > Under system variables > create a new variable "GOPATH", value "C:\Go\bin"
  * For Mac user: open terminal, type export GOPATH=$GOROOT/bin
4. Project location of "go-angular"
  * For windows user, extract this project to C:\Go\src\
  * For Mac user, extract this project to anywhere you want

Install: Project dependencies
=======
3. Open command line/terminal: navigate to frontend folder, type "npm install"
4. Open command line/terminal: navigate to backend folder, type "go get"

Run it
======
5. Front end
   * Open command line/terminal: navigate to frontend folder, type "npm start"
   * Open browser: http://localhost:8080/
6. Back end
   * Open command line/terminal: navigate to backend folder, type "go run main.go"
   * Open browser: http://localhost:8000/friend/


**in case you hit any $GOPATH error while running go command, please follow instruction https://golang.org/doc/install
