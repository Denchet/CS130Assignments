:: _runGoApp.bat
::  by Allen J. Mills
::  1 / 27 / 2015

:: note:
:: place this file in a file at or near your .go files

echo off
echo Run Go App

:: Update Path ---------------------------------------------------------
:: There are three requirements to run goApp
::		- GoLang
::		- GoApp
::		- Python 2.7
:: Ensure the paths below are to your copy of these folders
PATH=%path%;C:\Go
PATH=%path%;C:\Users\Dan-Null\goappengine\go_appengine
PATH=%path%;C:\Python27

:: Open A CMD with new Path -----------------------------------
echo Open a browser window to http://localhost:8080/
echo. 
echo. 
cmd 
exit