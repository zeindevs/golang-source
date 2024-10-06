#include <iostream>
#include <thread>

#if defined(_WIN32)
#include <windows.h>
typedef void(__stdcall *CallGoFunction)();
#else
#include <dlfcn.h>
typedef void (*CallGoFunction)();
#endif //_WIN32

void runGoFunctionInThread(CallGoFunction callGoFunc)
{
    std::cout << "Calling Go function in background..." << std::endl;
    callGoFunc();
}

int main()
{
    CallGoFunction callGoFunc = nullptr;

#if defined(_WIN32)
    HMODULE hDll = LoadLibrary("libmain.dll");
    if (!hDll)
    {
        std::cerr << "Failed to load DLL!" << std::endl;
        return 1;
    }

    callGoFunc = (CallGoFunction)GetProcAddress(hDll, "CallGoFunction");
    if (!callGoFunc)
    {
        std::cerr << "Failed to find CallGoFunction in DLL!" << std::endl;
        FreeLibrary(hDll);
        return 1;
    }
#else
    void *handle = dlopen("libmain.so", RTLD_LAZY);
    if (!handle)
    {
        std::cerr << "Failed to load shared library: " << dlerror() << std::endl;
        return 1;
    }

    callGoFunc = (CallGoFunction)dlsym(handle, "CallGoFunction");
    const char *dlsym_error = dlerror();
    if (dlsym_error)
    {
        std::cerr << "Failed to find CallGoFunction: " << dlsym_error << std::endl;
        dlclose(handle);
        return 1;
    }
#endif

    std::thread goThread(runGoFunctionInThread, callGoFunc);
    goThread.join();

    std::cout << "Go function execution finished!" << std::endl;

#if defined(_WIN32)
    FreeLibrary(hDll);
#else
    dlclose(handle);
#endif // _WIN32

    return 0;
}
