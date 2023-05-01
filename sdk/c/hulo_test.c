#include <stdio.h>
#include "libhulo/hulo_build.h"

CallRequest sayHello(CallRequest req)
{
    callRequest request = callRequestBuild(req);
    printf("Func: %d, Arg: %d", request.func, request.arg);
    return CallResponseBuild({.buf = "I'm C"});
}

int main(int argc, char **argv)
{
    HuloBuilder();
    HuloCall("SayHello", sayHello);
    HuloRun();
    return 0;
}