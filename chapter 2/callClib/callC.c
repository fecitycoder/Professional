#include <stdio.h>
#include "callC.h"

void cHello(){
    printf("Hello fromC!\n");
}

void printMessage(char* message){
    printf("Go send me %s\n", message);
}