# HW1

Сервис браниновария квартир, вся недвижимость находится в управлении сервиса.

##  Tермины:
- Unit – Недвижимость доступная для бронирования.
- Booking – Процесс резервирования unit на указанные даты.
- Availability – Статус доступности unit в определённые даты.
- Payment – Процесс оплаты за бронирование.
- User – Пользователь.
- Admin – Администратор платформы, имеет право изменять unit'ы.

## События:
1. Поиск:
   - Юзер предоставляется список юнитов для указанных дат.
   - Юзер может открыть карточку юнита с подробной информацией.

2.  Бронирование:
   - Пользователь выбирает юнит и даты.
   - Пользователь оплачивает стоимость аренды.

3. Управление бронированиями:
   - Пользователь может просматривать свои активные и прошедшие бронирования.
   - Отмена бронирования.

4. Администрирование квартиры:
   - Админ добавляет/удаляет/изменяет юнит.

## Cущности:
- Unit:
  - Адрес, описание, характеристики, фотографии.

- Booking:
  - Информация о бронировании, платеже, пользователе и юните.

- User:
  - Имя, контакты, доступ (guest, admin).

- Payment:
  - Cвязь с бронированием, сумма, статус (в ожидании/оплачено/отменено).

## Агрегаты:
- UnitAggregate: управляет состоянием квартиры (включая её бронирования).
- BookingAggregate: включает детали бронирования, статус и связанные платежи.

## Объекты значений:
- DateRange - диапазон дат.
- Address - город, улица, номер дома, номер квартиры.
