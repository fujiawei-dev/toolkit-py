{{SLASH_COMMENTS}}

#ifndef {{APP_NAME_UPPER}}__DEBUG_TOOLS_CRASH_STACK_H
#define {{APP_NAME_UPPER}}__DEBUG_TOOLS_CRASH_STACK_H

#include <QString>
#include <windows.h>

typedef struct STACK {
    STACK *ebp;
    PBYTE returnAddress;
    DWORD params[0];
} STACK, *pSTACK;

class CrashStack {
private:
    PEXCEPTION_POINTERS m_pException;

private:
    static QString GetModuleByReturnAddress(const PBYTE &returnAddress, PBYTE &moduleAddress);
    QString GetCallStack(PEXCEPTION_POINTERS pException);
    static QString GetVersionStr();

public:
    explicit CrashStack(PEXCEPTION_POINTERS pException);

    QString GetExceptionInfo();
};

long __stdcall CrashHandler(_EXCEPTION_POINTERS *pException);

#endif// CRASH_STACK_H
