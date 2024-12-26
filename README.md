# Calculator Service

Содержание
- [О проекте](#Опроекте)
- [Поддерживаемые функции](#ПоддерживаемыеФункции)
- [Как запустить](#КакЗапустить)


# О проекте
  Проект выполняет базовые вычисления. Просто передайте ей строку с математическим выражением и получите ответ!

# Поддерживаемые функции
Сервис калькулятор поддерживает базовые математические операции: +, -, *, / (сложение, вычитание, умножение и деление)

А также операторы приоретизации: ( и )

# Как запустить
## Устанавливаем Git
- __Windows__. Скачайте [Git для Windows](https://git-scm.com/download/win), запустите exe-файл, следуйте инструкциям.
- __macOS__. Скачайте [Git для macOS](https://sourceforge.net/projects/git-osx-installer/) и запустите dmg-файл. Если он не запускается, зайдите в Системные настройки — Безопасность и нажмите кнопку Open anyway (Всё равно открыть).
- __Linux__. Установите Git через встроенный менеджер пакетов. Если у вас Ubuntu, используйте команду sudo apt-get install git. Команды для других дистрибутивов можно посмотреть [здесь](https://git-scm.com/download/linux).

## Устанавливаем Go
Проект написан на языке Go(GoLang), для запуска сервера вам необходимо будет установить его:

Перейдите на официальный сайт и следуйте инструкциям: [сслыка](https://go.dev/doc/install)

## Скачиваем репозиторий
Откройте командную консоль и вставьте команду:

`git clone https://github.com/SanyaDr/FirstSprintExam.git`

## Запуск сервера
> [!Caution]
> Если для тестирования вы используете систему Windows, настоятельно рекомендую использовать консоль PowerShell, так как возникают проблемы с кавычками у стандартной командной строки Windows

Далее, вам необходимо открыть в консоли место расположения проекта. 

В моем случае это: `C:\Users\User\FirstSprintExam`

Затем перейдите в каталог _cmd_:  `cd cmd`

И наконец, запустите сервер командой:

`go run main.go`

Если всё прошло успешно, вы увидите в консоли время запуска и порт на котором запущен сервер. По умолчанию сервер запускается на порту _8080_

# Использование калькулятора

К калькулятору вы можете обратиться по адресу `/api/v1/calculate`

Однако, если вы попробуете открыть его через обычный браузер `http://localhost:8080/api/v1/calculate`, то вы получите ответ с текстом ___Method is not allowed___

Сервер принимает только POST запросы. Вот варианты того как вы можете их отправить:

### Через консоль
Использование _curl_

Для отправки запроса введите в консоль команду:

- Для стандартной консоли Windows:

`curl localhost:8080/api/v1/calculate --header Content-Type:application/json --data {"""expression""":"""2+2"""}`
- Для PowerShell и других:

`curl 'localhost:8080/api/v1/calculate' --data '{"expression": "2+2"}'`

Вместо _2+2_ впишите своё выражение которое хотите посчитать

### С помощью Postman
Postman - Сервис для создания, тестирования, документирования, публикации и обслуживания API.

Скачать его вы сможете на [официальном сайте](https://www.postman.com/). Для тестирования вам потребуется скачать Postman Agent, так как тестировать вы будете на локальной машине. 

#### Как пользоваться Postman
![image](https://github.com/user-attachments/assets/a018e5be-a09f-4684-99b0-4e68b6af517d)
Вы видите главный экран Postman. В верхней части вы увидете вкладки, как в браузере. Создайте новую вкладку.

![image](https://github.com/user-attachments/assets/597479bd-c2aa-4f88-b1f0-c7a06e07fd8a)
Заполните необходимые поля, как показано на скриншоте.
- Выбран POST запрос
- В адресной ссылке введено: `http://localhost:8080/api/v1/calculate`
- Во вкладке _Body_ выбран режим _raw_ типа JSON
- В текстовом редакторе написано код формата:
  `
  {
    "expression": "2+2*2"
  }
   `

  # Возвращаемые ошибки и коды
  200 - Ответ успешно получен, ошибок нет
  
  422 - Ошибка пользователя, например выражение введено с ошибкой

  500 - Ошибка на стороне сервера

