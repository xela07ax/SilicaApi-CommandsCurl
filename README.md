## Testing apps
I have curl app from testing api 

## Create project
```sh
curl -H "Content-Type: application/json" -X POST http://localhost:30339/ngapi/createProject -d "{\"Name\":\"Fuels\",\"Method\":\"Maven\"}"
```
Ожидаем ответ
```json
{"FuncName":"CreateTest","Text":"Проект создан","Status":0,"List":null,"Show":true,"UpdNum":0}
```
## Read all projects
Смотрим на все созданные проекты
```sh
curl  -X POST http://localhost:30339/ngapi/getAllProjects
```
Ожидаем ответ
```json
[{"Id":8000001,"Name":"Fuels","Status":0,"Lang":"","Method":"Maven","CreateDate":"2020-10-19T18:26:15.1521269+03:00","UpdateDate":"2020-10-19T18:26:15.1521269+03:00","EndDate":"1754-08-31T01:43:41.128654848+03:00","Playlists":{},"PlaylistArray":null}]
```

## Import from Git project
Заимпортируем открытый проект с гита
```sh
curl -H "Content-Type: application/json" -X POST http://localhost:30339/ngapi/gitClone -d "{\"ProjectId\":8000001,\"Url\":\"https://github.com/xela07ax/SilicaNG_1.git\"}"
```
Ожидаем ответ
```json
{"FuncName":"CreateTest","Text":"Проект успешно склонирован с гита","Status":0,"List":null,"Show":true,"UpdNum":0}
```
## Add command CMD runner from the project
Добавим команду для запуска
```sh
curl -H "Content-Type: application/json" -X POST http://localhost:30339/ngapi/addCmd -d "{\"ProjectId\":8000001,\"Command\":\"mvn\",\"Args\":[\"test\"]}"
```
Ожидаем ответ
```json
{"FuncName":"CreateTest","Text":"Команда CMD успешно добавлена в проект","Status":0,"List":null,"Show":true,"UpdNum":0}
```
## Get Cmds is project
```sh
curl -H "Content-Type: application/json" -X POST http://localhost:30339/ngapi/getCmdsFromProject -d "{\"ProjectId\":8000001}"
```
Ожидаем ответ
```json
[{"Id":1,"Command":"mvn","Args":["test"]},{"Id":2,"Command":"mvn","Args":["test"]},{"Id":3,"Command":"mvn","Args":["test"]}]
```
## Remove Cmd from project !!*not check
```sh
curl -H "Content-Type: application/json" -X POST http://localhost:30339/ngapi/delCmdFromProject -d "{\"ProjectId\":8000001,\"Command\":{\"Id\":2}}"
```
Ожидаем ответ
```json
{"FuncName":"CreateTest","Text":"CMD команда удалена успешно","Status":0,"List":null,"Show":true,"UpdNum":0}
```
## Start Cmd
```sh
curl -H "Content-Type: application/json" -X POST http://localhost:30339/ngapi/run -d "{\"ProjectId\":8000001,\"CmdId\":2}"
```
Ожидаем ответ
```json
{"FuncName":"CreateTest","Text":"CMD команда удалена успешно","Status":0,"List":null,"Show":true,"UpdNum":0}
```
