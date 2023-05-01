#ifndef __HULO_LIB_H__
#define __HULO_LIB_H__

typedef char *String;
typedef String error;
typedef String (*Call)(String);
typedef char Boolean;
typedef String CallResponse;
typedef String CallRequest;
typedef struct
{
    void *ptr;
} Hulo;

String huloCall(Call, String);

typedef struct
{
    String buf;
    Boolean ok;
} callResponse;

typedef struct
{
    String func;
    String arg;
} callRequest;

callRequest callRequestBuild(CallRequest);
CallResponse callResponseBuild(callResponse);

#define CallResponseBuild(req) callResponseBuild((callResponse)req)
#endif