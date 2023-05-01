#include "libhulo.h"

#define HuloBuilder() Hulo *hulo = newHulo()
#define HuloCall(name, callback) registerCall(hulo, name, callback)
#define HuloRun() run(hulo, argv[2])