package main

/**
 * @author : ${USER}
 * @created : ${DATE}
**/

import (
    "log"
    "os"
)

func main() {

//Catch panic like this, so it shows with full stacktrace in logs
defer func() {
    if r := recover(); r != nil {
        log.Printf("caught panic: %v", r)
        os.Exit(1)
    }
}()

//By running otherwise known "main" inside separate function you can control which errors,
//should stop the program all together, and which ones are only going to be handled
//inside this dedicated run function
if err := run(); err != nil {
    log.Fatalf("fatal exception: %v", err)
}
}

//This is the dedicated run function which hosts your usual main code and controls what,
//will be logged here and what is passed as fatal error further
func run() error {
    //Code here

return nil
}
