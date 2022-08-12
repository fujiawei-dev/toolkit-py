import multiprocessing
import os

import servicemanager
import win32event
import win32service
import win32serviceutil

from {{project_slug.snake_case}}.logger import logging
from {{project_slug.snake_case}}.__main__ import main

log = logging.getLogger(__name__)

class AppWindowsService(win32serviceutil.ServiceFramework):
    _svc_name_ = {{project_slug.snake_case}}
    _svc_display_name_ = {{project_slug.words_capitalized}}
    _svc_description_ = {{ project_short_description }}

    def __init__(self, args):
        win32serviceutil.ServiceFramework.__init__(self, args)
        self.hWaitStop = win32event.CreateEvent(None, 0, 0, None)

    def SvcStop(self):
        self.ReportServiceStatus(win32service.SERVICE_STOP_PENDING)
        win32event.SetEvent(self.hWaitStop)

    def SvcDoRun(self):
        servicemanager.LogMsg(
            servicemanager.EVENTLOG_INFORMATION_TYPE,
            servicemanager.PYS_SERVICE_STARTED,
            (self._svc_name_, ""),
        )

        # determine if application is a script file or frozen exe
        if getattr(sys, "frozen", False):
            application_path = os.path.dirname(sys.executable)
        else:
            application_path = os.path.dirname(__file__)

        process = multiprocessing.Process(target=main)
        process.start()  # start process

        while True:
            rc = win32event.WaitForSingleObject(self.hWaitStop, 5000)
            # The service is stopping
            if rc == win32event.WAIT_OBJECT_0:
                process.kill()
                break


if __name__ == "__main__":
    import sys
    from multiprocessing import freeze_support

    # https://github.com/pyinstaller/pyinstaller/issues/2023
    freeze_support()

    if len(sys.argv) == 1:
        servicemanager.Initialize()
        servicemanager.PrepareToHostSingle(AppWindowsService)
        servicemanager.StartServiceCtrlDispatcher()
    else:
        win32serviceutil.HandleCommandLine(AppWindowsService)
