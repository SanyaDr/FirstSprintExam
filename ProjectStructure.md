# Структура проекта
_Описание структуры проекта_

├── FirstSprintExam[^1] </br>
├── cmd[^2]</br>
│ $~~~~~~~~$ └─ main.go </br>
├── internal[^3]</br>
│ $~~~~~~~~$ ├─ app[^4] </br>
│ $~~~~~~~~~$ │ $~~~~~~$ ├ app.go </br>
│ $~~~~~~~~~$ │ $~~~~~~$ └ handlers.go </br>
│ $~~~~~~~~$ └─ transport [^5] </br>
│ $~~~~~~~~~~~~~~~~~~~~~$ ├ getExpressionForCalculator.go </br>
│ $~~~~~~~~~~~~~~~~~~~~~$ └ returnJson.go </br>
├── pkg[^6]</br>
│ $~~~~~~~~$ └─ calculation </br>
│ $~~~~~~~~~~~~~~~~~$ ├─ errors.go </br>
│ $~~~~~~~~~~~~~~~~~$ ├─ separators.go </br>
│ $~~~~~~~~~~~~~~~~~$ ├─ simplifyer.go </br>
│ $~~~~~~~~~~~~~~~~~$ ├─ simpleCalc.go </br>
│ $~~~~~~~~~~~~~~~~~$ └─ calculator.go </br>
│</br>
└── README.md


[^1]: Корневая папка проекта
[^2]: _cmd_ - Содержит _main_ файл, точка входа в приложение
[^3]: _internal_ - Содержит код, который используется исключильно внутри проекта. Запрещено к импорту в иные проекты
[^4]: _app_ - Точка запуска приложения. Содержит run-методы, которые вызываются из _cmd_
[^5]: _transport_ - Здесь храним http-настройки сервера, хендлеры, порты и так далее.
[^6]: _pkg_ - Код, который можно экспортировать в другие проекты
