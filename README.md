# Агрегатор цен игр с разных онлайн площадок
Примеры запрос/ответ API

`/getname?name=Saints Row`

endpoint выводит 10 наиболее похожих имен к исходному запросу.
```
[
  "Saints Row 2",
  "Saints Row IV - Saints Row IV Season Pass",
  "Saints Row IV",
  "Saints Row IV - The Super Saints Pack",
  "Saints Row IV - How the Saints Save Christmas",
  "Saints Row IV - Game On",
  "Saints Row IV - GAT V Pack",
  "Saints Row IV - Anime Pack",
  "Saints Row IV - College Daze",
  "Saints Row IV - The Rectifier"
]
```

`/getname?name=Dota 2`

Если есть точное совпадение - результат одно значение
```
[
  "Dota 2"
]
```


`/compareprice?name=Saints Row 2`

Отправляет json файлы с необходимыми данными об игре
Есть случаи, где игра есть в магазине и где игры нет.
ЭТОТ ЭНДПОИНТ РАБОТАЕТ ТОЛЬКО С КОРРЕКТНЫМ ИМЕНЕМ
т.е. сначала идет /getname, а потом уже /compareprice
```
[
  {
    "store_name": "steam",
    "store_app_id": 9480,
    "store_app_name": "Saints Row 2",
    "store_price": "62 руб.",
    "store_image": "https://cdn.akamai.steamstatic.com/steam/apps/9480/header.jpg?t=1620659051",
    "store_app_url": "https://store.steampowered.com/app/9480"
  },
  {
    "store_name": "steampay",
    "store_app_name": "Saints Row 2",
    "status": "game not found in store"
  },
  {
    "store_name": "gog",
    "store_app_id": 1430740458,
    "store_app_name": "Saints Row 2",
    "store_price": "69 руб.",
    "store_image": "https://images-3.gog-statics.com/e5054aacbe4d66cc91783dfe5d2eb996e49d08523f8ed6f2cb07bd9cc747aed2",
    "store_app_url": "https://gog.com/game/saints_row_2"
  },
  {
    "store_name": "platiru",
    "store_app_id": 3045958,
    "store_app_name": "Saints Row 2",
    "store_price": "239 руб.",
    "store_image": "https://graph.digiseller.ru/img.ashx?idd=3045958",
    "store_app_url": "https://www.plati.market/itm/3045958"
  }
]
```
Для связи интерфейса и api используются только два вышеупомянутых эндпоинта.
Также есть следующие эндпоинты:
`/steamprice`
`/steampayprice`
`/gogprice`
`/platiruprice`
