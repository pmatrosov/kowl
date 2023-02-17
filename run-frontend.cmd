@echo off
setlocal

for /f %%i in ('git.exe rev-parse --short^=8 HEAD') do set REACT_APP_CONSOLE_GIT_SHA=%%i
for /f %%i in ('git.exe rev-parse --abbrev-ref HEAD') do set REACT_APP_CONSOLE_GIT_REF=%%i
for /f %%i in ('git.exe log -1 --format^=%%ct HEAD') do set REACT_APP_BUILD_TIMESTAMP=%%i

echo GIT_SHA=%REACT_APP_CONSOLE_GIT_SHA%
echo GIT_REF=%REACT_APP_CONSOLE_GIT_REF%
echo GIT_TIME=%REACT_APP_BUILD_TIMESTAMP%

echo.
echo *** 1/3 BUILDING FRONTEND ***

set REACT_APP_ENABLED_FEATURES=REASSIGN_PARTITIONS

cd frontend
call npm install
call node ./scripts/start.js

endlocal
