url_git_node="git@bitbucket.org:cogtive/core.git"

echo "*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*"
echo "*-*-*-*-*-* Iniciando Deploy CORE *-*-*-*-*-*" 
echo "*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*"

if [ -d "core" ]; then
	rm -rf core
fi

mkdir core

if [ -d "core_temp" ]; then
	rm -rf core_temp
fi

mkdir core_temp
cd core_temp

git clone --depth=1 -b $1 $url_git_node .


echo "*-*-*-*-*-*-* Limpando  projeto *-*-*-*-*-*-*" 
dotnet clean 

echo "*-*-*-*-*-*-*   Publicando API   *-*-*-*-*-*-*"
cd Cogtive.Core.Api.V2

dotnet restore 
dotnet publish

cp -R bin/Debug/netcoreapp2.0/publish/ ../../core
mv ../../core/publish ../../core/Api

echo ""
echo "*-*-*-*-*-*-* Publicando API APLICACAO *-*-*-*-*-*-*"
cd ../Cogtive.Core.Api.Aplicacao

dotnet restore 
dotnet publish

cp -R bin/Debug/netcoreapp2.0/publish/ ../../core
mv ../../core/publish ../../core/Aplicacao

mkdir ../../core/Aplicacao/Util
cp -R Util/Arquivos/ ../../core/Aplicacao/Util/Arquivos
chmod 777 ../../core/Aplicacao/Util/Arquivos/Lote.xsd  ../../core/Aplicacao/Util/Arquivos/Pipa.xsd

echo ""
echo "*-*-*-*-*-*  Publicando  HEIMDALL  *-*-*-*-*-*"
cd ../Cogtive.Core.Heimdall

dotnet restore 
dotnet publish

cp -R bin/Debug/netcoreapp2.0/publish/ ../../core
mv ../../core/publish ../../core/Heimdall

echo ""
echo "*-*-*-*-*-*  Publicando  JAIMINHO  *-*-*-*-*-*"
cd ../Cogtive.Core.Jaiminho

dotnet restore 
dotnet publish

cp -R bin/Debug/netcoreapp2.0/publish/ ../../core
mv ../../core/publish ../../core/Jaiminho

echo ""
echo "*-*-*-*-*-*-*  Publicando VELMA  *-*-*-*-*-*-*"
cd ../Cogtive.Core.Velma

dotnet restore 
dotnet publish

cp -R bin/Debug/netcoreapp2.0/publish/ ../../core
mv ../../core/publish ../../core/Velma

for file in  `ls -1` ; do zip "$file".zip $file/*; done

cd ..
cd ..
rm -rf core_temp

echo "*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*"
echo "*-*-*-*-*-*-* Deploy finalizado *-*-*-*-*-*-*" 
echo "*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*"
	