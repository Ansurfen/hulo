#include "hulo.h"
#include "cJSON.h"

String huloCall(Call cb, String req)
{
    return cb(req);
}

CallResponse callResponseBuild(callResponse res)
{
    cJSON *root = cJSON_CreateObject();
    cJSON_AddStringToObject(root, "Buf", res.buf);
    cJSON_AddBoolToObject(root, "Ok", res.ok);
    return cJSON_PrintUnformatted(root);
}

callRequest callRequestBuild(CallRequest req)
{
    cJSON *root = cJSON_Parse(req);
    if (!root)
        return (callRequest){.func = "err", .arg = ""};
    callRequest ret;
    ret.func = cJSON_GetStringValue(cJSON_GetObjectItem(root, "Func"));
    ret.arg = cJSON_GetStringValue(cJSON_GetObjectItem(root, "Arg"));
    return ret;
}