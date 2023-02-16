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
call node ./scripts/build.js -- --profile

echo.
echo *** 2/3 COPYING FRONTEND ***

cd ..
del backend\pkg\embed\frontend\*.* /S /Q > nul
xcopy frontend\build backend\pkg\embed\frontend\ /E /Q

echo.
echo *** 3/3 BUILDING BACKEND ***

cd backend
go mod download
set CGO_ENABLED=0
go build -o ../build/console.exe ./cmd/api

echo.
echo DONE!
endlocal
pause
