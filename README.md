## counter-api

[herokuapp deployment](https://kasper-countere.herokuapp.com)

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

#### Персистентное хранение
В задаче предложено пользоваться персистентным хранилищем не выходя за рамки стандартной библиотеки, что исключает СУБД. Данные хранятся в gob-формате.