# go-lvl-four
1. Добавить в пример с файловым сервером возможность получить список всех файлов
на сервере (имя, расширение, размер в байтах)
2. С помощью query-параметра, реализовать фильтрацию выводимого списка по
расширению (то есть, выводить только .png файлы, или только .jpeg)
3. *Текущая реализация сервера не позволяет хранить несколько файлов с одинаковым
названием (т.к. они будут храниться в одной директории на диске). Подумайте, как
можно обойти это ограничение?
4. К коду, написанному в рамках заданий 1-3, добавьте тесты с использованием
библиотеки httptest.
