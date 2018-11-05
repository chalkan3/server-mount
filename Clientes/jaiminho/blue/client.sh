url_git_node="git@bitbucket.org:cogtive/client.git"

url_api_csharp="https://jaiminho.cogtive.com.br/api/v2"
url_api_aplicacao="https://jaiminho.cogtive.com.br/api"

port_api="123"

url_api_aplicacao_valida=$(echo $url_api_aplicacao | sed -r 's/\//\\\//g')
url_api_csharp_valida=$(echo $url_api_csharp | sed -r 's/\//\\\//g')

echo "*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*"
echo "*-*-*-*-* Iniciando  Deploy  Client *-*-*-*-*" 
echo "*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*"

if [ -d "client_temp" ]; then
	rm -rf client_temp
fi

mkdir client_temp
cd client_temp

git clone --depth=1 -b $1 $url_git_node .

cd src/common

sed -i.bak -r -e 's/var[[:blank:]]*APIAplicacao[[:blank:]]*\=[[:blank:]]*\"[^\"]*\"/var APIAplicacao = \"'$url_api_aplicacao_valida'\"/' \
				-e 's/var[[:blank:]]*APIV2[[:blank:]]*\=[[:blank:]]*\"[^\"]*\"/var APIV2 = \"'$url_api_csharp_valida'\"/' Constants.js

cd ../../
sed -i.bak -r -e 's/var[[:blank:]]*Porta[[:blank:]]*\=[[:blank:]]*\"[^\"]*\"/var Porta = '$port_api'/' producao.server.js


echo "*-*-*-*-* Restaurando Dependências  *-*-*-*-*"
yarn install

echo "*-*-*-*-*-*-*-*-* Copilando *-*-*-*-*-*-*-*-*"
yarn deploy

cd js
rm -rf assets

cp -R ../src/assets/ ./

cd  ../../

if [ -d "client" ]; then
	rm -rf client
fi

cp -R client_temp/js ./
mv js client

cp -R client_temp/producao.server.js client/server.js

cd client
npm install express

cd ..
rm -rf client_temp

echo "*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*"
echo "*-*-*-*-*-*-* Deploy finalizado *-*-*-*-*-*-*" 
echo "*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*"
	