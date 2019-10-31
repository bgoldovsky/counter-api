## counter-api

Используя только средства стандартной библиотеки, разработать HTTP-сервер. Сервер должен:

* На каждый запрос возвращать значение счетчика, считающего общее количество обработанных сервером запросов за последние 60 секунд;
* Продолжать возвращать корректное значение счетчика после перезапуска приложения используя персистентное хранилище.

#### Ожидается:

* Аккуратный, хорошо-структурированный код с применением устоявшихся т.н. best practices
* Документирование кода в местах на усмотрение автора

#### Будет плюсом:

* Тесты
* Готовый, к деплою на production-стенд, код
* Использование Docker
* Комментарии, почему было выбрано то или иное решение

## Комментарии

У меня не слишком много опыта работы с Go, поэтому ряд решений может оказаться спорным или ошибочным.
Будет здорово получить по ним фидбек и поговорить про объектную модель в Go, в том числе про применимость к ней практик из ООП-мира вроде DDD и Event Sourcing.

#### Глобальный обработчик 
В метод http.ListenAndServe() передается структура для обработки всех запросов к сервису. Насколько мне известно, это далеко не best practices и если бы задание не указывало на единый endpoint сервиса, то я предпочел бы использовать http.HandleFunc()

#### Инициализация пакетов.
После C#/Java и конструкторов не совсем понятно, как правильно инициализации пакеты, если туда необходимо передать аргументы из main.
Я использовал функцию Init(someArgsHere). Для приложения это выглядит нормально (хотя и режет глаз), при работе над библиотекой очередность вызова методов стала бы неочевидна для пользователей.
init(), к сожалению, для данных целей не очень подходит

#### Персистентное хранение
В задачке предложено пользоваться персистентным хранилищем не выходя за рамки стандартной библиотеки, что сразу исключило СУБД. Мне показалось оптимальным хранить данные в gob-формате.
Если бы можно было обойтись in-memory хранилищем, то вместо дорогостоящей блокировки sync.Mutex был бы использован lock-free пакет atomic для cas-инкремента. 

#### panic/recover
В одном месте я использую на выходе из сервисного слоя функцию recover().
Не знаю насколько это применимо в Go, но мне показалось логичным сделать что-то вроде global event handler из мира C#/Java на тот случай, если какая-то из библиотечных функций кинет panic.