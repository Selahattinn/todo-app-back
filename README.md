# todo-app-back

```
├─ bin                           //The folder where the binary file was created
├─ cmd                           //The code that started it all
├─ config.yml                    //Config showing which information the backend server will run
├─ go.mod                        //3rd party libraries
├─ go.sum                        //Sums and versions of 3rd party libraries
├─ makefile                      //MakeFile for version control ,creation of binary file tests 
└─ pkg                           //Server codes 
   ├─ api                        //Api Layer for all aplication
   ├─ model                      //Models for every type of object
   ├─ repository                 //DB Layer
   │  ├─ todo
   ├─ server                     //Server Layer for all aplication. Main part of http server
   ├─ service                    //Service Layer
   │  ├─ todo
   ├─ static                     //Auto generated static files
   └─ version                    //Version control&save for git