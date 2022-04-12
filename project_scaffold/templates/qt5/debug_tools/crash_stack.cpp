{{SLASH_COMMENTS}}

#include "crash_stack.h"
#include <QDateTime>
#include <QDir>
#include <QTextCodec>
#include <Wbemidl.h>
#include <comdef.h>
#include <cstdio>
#include <dbghelp.h>
#include <qdebug.h>
#include <tlhelp32.h>

CrashStack::CrashStack(PEXCEPTION_POINTERS pException) {
    m_pException = pException;
}

QString CrashStack::GetModuleByReturnAddress(const PBYTE &returnAddress, PBYTE &moduleAddress) {
    HANDLE handleSnapshot;
    wchar_t moduleName[MAX_PATH] = {0};
    MODULEENTRY32 moduleEntry = {sizeof(moduleEntry)};

    handleSnapshot = CreateToolhelp32Snapshot(TH32CS_SNAPMODULE, 0);
    if ((handleSnapshot != INVALID_HANDLE_VALUE) &&
        Module32First(handleSnapshot, &moduleEntry)) {
        do {
            if (DWORD(returnAddress - moduleEntry.modBaseAddr) < moduleEntry.modBaseSize) {
                lstrcpyn((char *) moduleName, moduleEntry.szExePath, MAX_PATH);
                moduleAddress = moduleEntry.modBaseAddr;
                break;
            }
        } while (Module32Next(handleSnapshot, &moduleEntry));
    }

    CloseHandle(handleSnapshot);
    return QByteArray((char *) moduleName);
}

QString CrashStack::GetCallStack(PEXCEPTION_POINTERS pException) {
    PBYTE moduleAddress;
    char buffer[256] = {0};
    QString result;

    STACK Stack = {nullptr, nullptr};
    pSTACK Ebp;

    if (pException)//fake frame for exception address
    {
        Stack.ebp = (pSTACK) pException->ContextRecord->Ebp;
        Stack.returnAddress = (PBYTE) pException->ExceptionRecord->ExceptionAddress;
        Ebp = &Stack;
    } else {
        Ebp = (pSTACK) &pException - 1;//frame address of GetCallStack()

        // Skip frame of GetCallStack().
        if (!IsBadReadPtr(Ebp, sizeof(pSTACK)))
            Ebp = Ebp->ebp;//caller ebp
    }

    // Break trace on wrong stack frame.
    for (; !IsBadReadPtr(Ebp, sizeof(pSTACK)) && !IsBadCodePtr(FARPROC(Ebp->returnAddress));
         Ebp = Ebp->ebp) {
        // If module with Ebp->returnAddress found.
        memset(buffer, 0, sizeof(0));
        sprintf(buffer, "\n%08X  ", (unsigned int) Ebp->returnAddress);
        result.append(buffer);

        QString moduleName = this->GetModuleByReturnAddress(Ebp->returnAddress, moduleAddress);
        if (moduleName.length() > 0) {
            result.append(moduleName);
        }
    }

    return result;
}

QString CrashStack::GetVersionStr() {
    OSVERSIONINFOEX V = {sizeof(OSVERSIONINFOEX)};//EX for NT 5.0 and later

    if (!GetVersionEx((POSVERSIONINFO) &V)) {
        ZeroMemory(&V, sizeof(V));
        V.dwOSVersionInfoSize = sizeof(OSVERSIONINFO);
        GetVersionEx((POSVERSIONINFO) &V);
    }

    if (V.dwPlatformId != VER_PLATFORM_WIN32_NT)
        V.dwBuildNumber = LOWORD(V.dwBuildNumber);//for 9x HIWORD(dwBuildNumber) = 0x04xx

    return QString("Windows:  %1.%2.%3, SP %4.%5, Product Type %6\n")
            .arg(V.dwMajorVersion)
            .arg(V.dwMinorVersion)
            .arg(V.dwBuildNumber)
            .arg(V.wServicePackMajor)
            .arg(V.wServicePackMinor)
            .arg(V.wProductType);
}

QString CrashStack::GetExceptionInfo() {
    WCHAR moduleName[MAX_PATH];
    PBYTE moduleAddress;

    QString result;
    char buffer[512] = {0};

    QString version = GetVersionStr();
    result.append(version);

    result.append("Process:  ");
    GetModuleFileName(nullptr, (char *) moduleName, MAX_PATH);

    //    result.append(QString::fromWCharArray(moduleName));
    result.append(QByteArray::fromRawData((char *) moduleName, MAX_PATH));
    result.append("\n");

    // If exception occurred.
    if (m_pException) {
        EXCEPTION_RECORD &E = *m_pException->ExceptionRecord;
        CONTEXT &C = *m_pException->ContextRecord;

        memset(buffer, 0, sizeof(buffer));
        sprintf(buffer, "Exception Address:  %08X  ", (int) E.ExceptionAddress);
        result.append(buffer);
        QString module = GetModuleByReturnAddress((PBYTE) E.ExceptionAddress, moduleAddress);
        if (module.length() > 0) {
            result.append(" \nModule: ");
            result.append(module);
        }

        memset(buffer, 0, sizeof(buffer));
        sprintf(buffer, "\nException Code:  %08X\n", (int) E.ExceptionCode);
        result.append(buffer);

        if (E.ExceptionCode == EXCEPTION_ACCESS_VIOLATION) {
            // Access violation type - Write/Read.
            memset(buffer, 0, sizeof(buffer));
            sprintf(buffer, "%s Address:  %08X\n",
                    (E.ExceptionInformation[0]) ? "Write" : "Read", (int) E.ExceptionInformation[1]);
            result.append(buffer);
        }


        result.append("Instruction: ");
        for (int i = 0; i < 16; i++) {
            memset(buffer, 0, sizeof(buffer));
            sprintf(buffer, " %02X", PBYTE(E.ExceptionAddress)[i]);
            result.append(buffer);
        }

        result.append("\nRegisters: ");

        memset(buffer, 0, sizeof(buffer));
        sprintf(buffer, "\nEAX: %08X  EBX: %08X  ECX: %08X  EDX: %08X", (unsigned int) C.Eax, (unsigned int) C.Ebx, (unsigned int) C.Ecx, (unsigned int) C.Edx);
        result.append(buffer);

        memset(buffer, 0, sizeof(buffer));
        sprintf(buffer, "\nESI: %08X  EDI: %08X  ESP: %08X  EBP: %08X", (unsigned int) C.Esi, (unsigned int) C.Edi, (unsigned int) C.Esp, (unsigned int) C.Ebp);
        result.append(buffer);

        memset(buffer, 0, sizeof(buffer));
        sprintf(buffer, "\nEIP: %08X  EFlags: %08X", (unsigned int) C.Eip, (unsigned int) C.EFlags);
        result.append(buffer);
    }

    result.append("\nCall Stack:");
    QString sCallstack = this->GetCallStack(m_pException);
    result.append(sCallstack);

    return result;
}

#ifdef Q_OS_WIN
long __stdcall CrashHandler(_EXCEPTION_POINTERS *pException) {
    CrashStack crashStack(pException);
    QString sCrashInfo = crashStack.GetExceptionInfo();
    QString file_path = QDir::currentPath();

    QDir *folder_path = new QDir;
    bool exist = folder_path->exists(file_path.append("/debug"));
    if (!exist) {
        folder_path->mkdir(file_path);
    }
    delete folder_path;

    QString sFileName = file_path + "/crash.log";

    QFile file(sFileName);
    if (file.open(QIODevice::WriteOnly | QIODevice::Truncate)) {
        file.write(sCrashInfo.toUtf8());
        file.close();
    }

    return EXCEPTION_EXECUTE_HANDLER;
}
#else
long CrashHandler(EXCEPTION_POINTERS *pException) {
    {
        QDir *dumpDir = new QDir;
        bool exist = dumpDir->exists("../dump/");
        if (!exist) dumpDir->mkdir("../dump/");
    }
    QDateTime current_date_time = QDateTime::currentDateTime();
    QString current_date = current_date_time.toString("yyyy_MM_dd_hh_mm_ss");
    QString time = current_date + ".dump";
    EXCEPTION_RECORD *record = pException->ExceptionRecord;
    QString errCode(QString::number(record->ExceptionCode, 16));
    QString errAddress(QString::number((uint) record->ExceptionAddress, 16));
    QString errFlag(QString::number(record->ExceptionFlags, 16));
    QString errParams(QString::number(record->NumberParameters, 16));

    qDebug() << "errCode: " << errCode;
    qDebug() << "errAddress: " << errAddress;
    qDebug() << "errFlag: " << errFlag;
    qDebug() << "errParams: " << errParams;

    HANDLE hDumpFile = CreateFile(
            reinterpret_cast<LPCSTR>(QString("../dump/" + time).utf16()),
            GENERIC_WRITE,
            0,
            nullptr,
            CREATE_ALWAYS,
            FILE_ATTRIBUTE_NORMAL,
            nullptr);

    if (hDumpFile != INVALID_HANDLE_VALUE) {
        MINIDUMP_EXCEPTION_INFORMATION dumpInfo;
        dumpInfo.ExceptionPointers = pException;
        dumpInfo.ThreadId = GetCurrentThreadId();
        dumpInfo.ClientPointers = TRUE;
        MiniDumpWriteDump(
                GetCurrentProcess(),
                GetCurrentProcessId(),
                hDumpFile,
                MiniDumpNormal,
                &dumpInfo,
                nullptr,
                nullptr);
        CloseHandle(hDumpFile);
    } else {
        qDebug() << "hDumpFile == null";
    }
    return EXCEPTION_EXECUTE_HANDLER;
}
#endif
