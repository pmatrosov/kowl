@echo off
setlocal

for /f %%i in ('git.exe rev-parse --short^=8 HEAD') do set GIT_SHA=%%i
for /f %%i in ('git.exe rev-parse --abbrev-ref HEAD') do set GIT_REF=%%i
for /f %%i in ('git.exe log -1 --format^=%%ct HEAD') do set GIT_TIME=%%i

echo GIT_SHA=%GIT_SHA%
echo GIT_REF=%GIT_REF%
echo GIT_TIME=%GIT_TIME%

docker build . --build-arg GIT_SHA=%GIT_SHA% --build-arg GIT_REF=%GIT_REF% --build-arg GIT_TIME=%GIT_TIME% --no-cache --tag pmatrosov/kowl:latest
endlocal
pause