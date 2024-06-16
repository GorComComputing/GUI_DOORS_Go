# Операционная система Doors в браузере

![img/Screenshot5.png](img/Screenshot5.png)
![img/Screenshot4.png](img/Screenshot4.png)

Оконный графический интерфейс операционной системы Doors. Версия для WebAssembly, запускается в браузере.

Содержит:
- базовую графическую библиотеку
- невытесняющий планировщик задач
- набор графических приложений

Из операционной системы Doors можно управлять Умным домом и простматривать IP-камеры.

Написан на Go для WebAssembly.

Статус проекта: Разрабатывается.


### Приложение состоит из двух частей:
- Back-end  
	Это серверная часть, запускается на устройстве  
	Написан на Go версии 1.18.1.  
	Файлы *.go лежат в каталоге ./Server  
	Файл для сборки: ./Server/Makefile  
	Результат сборки: ./Server/DOORS  
- Front-end  
	Это клиентская часть, запускается в браузере  
	Написан на Go  
	Компилятор: TinyGo версия 0.28.1  
	Цель компиляции: Wasm  
	Файлы *.go лежат в корне каталога проекта ./  
	Результат сборки: ./Server/www/DOORS.wasm  
	Дополнительные библиотеки JavaScript:  
	- для связки JavaScript браузера с файлом DOORS.wasm в каталоге ./www/scripts/wasm  
	Стартовая страница ./www/index.html  
	
	Справочно: Компилятор TinyGo применен в проекте, потому что он создает бинарные файлы меньшего размера.  
			   Исходный код можно компилировать обычным компилятором Go с указанием цели компиляции Wasm.  

![img/Screenshot1.png](img/Screenshot1.png)
![img/Screenshot2.png](img/Screenshot2.png)
![img/Screenshot3.png](img/Screenshot3.png)

```
// Логин и пароль хранятся во внешнем файле secret.go в переменных:
var loginAdmin = "********"
var passwordAdmin = "********"
var loginUser = "********"
var passwordUser = "********"
```

```
# Для сборки и загрузки на сервер WASM
$ ./make.py wasm
# Для сборки и загрузки на сервер Web-сервера DOORS
$ ./make.py srv

# Остановить Web-сервер DOORS
$ ./make.py stop
# Запустить Web-сервер DOORS
$ ./make.py start
# Узнать запущен ли Web-сервер DOORS
$ ./make.py stat

# Для автоматческого добавление на GitHub
$ ./make.py git "Комментарий"

# Токен GitHub хранится во внешнем файле secret.py в переменной:
# tokenGit = "********"
```

2023-2024 Evgeny Goryachev    
Gor.Com 


