package models

import (
	"github.com/Vincent-Benedict/TPA-Web/database"
	"time"
)

type GameItem struct {
	ID              uint `gorm:"primary_key"`
	GameId          int
	ItemName        string
	ItemTransaction int
	ItemPrice       int
	ItemImage       string
	ItemDescription string `gorm:"default:'This Item can be applied to any weapon you own and can be scraped to look more worn. You can scrape the same sticker multiple times, making it a bit more worn each time, until it is removed from the weapon.'"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
}

func init() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.DropTableIfExists(&GameItem{})
	db.AutoMigrate(&GameItem{})
	SeedGameItem()
}

func SeedGameItem(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Create(&GameItem{
		GameId:          1,
		ItemName:        "2020 RMR Legends",
		ItemTransaction: 156,
		ItemPrice:       300000,
		ItemImage:       "https://community.akamai.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXU5A1PIYQNqhpOSV-fRPasw8rsQEl9Jg9SpIW1KgRr7OHFY28SvoyJl4iKm_vxPbzUhHgfuZEg2eyUpd2s0Qbsqhc4ZG33cI6VIFA5YgvS-wLrlb291JbtvJjMymwj5HcDj7YBnA/62fx62f",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          1,
		ItemName:        "2020 RMR Challengers",
		ItemTransaction: 172,
		ItemPrice:       300000,
		ItemImage:       "https://community.akamai.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXU5A1PIYQNqhpOSV-fRPasw8rsQEl9Jg9SpIW1KgRr7OHFY28SvoyJl4iKm_vxPbzUhHgfuZEg2eyUpd2s0Qbsqhc4ZG33cI6VIFA5YgvS-wLrlb291JbtvJjMymwj5HcDj7YBnA/62fx62f",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          1,
		ItemName:        "2020 RMR Contenders",
		ItemTransaction: 300,
		ItemPrice:       200000,
		ItemImage:       "https://community.akamai.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXU5A1PIYQNqhpOSV-fRPasw8rsQEl9Jg9SpIW1KgRr7OHFY28SvoyJl4iKm_vxPbzUhHgfuZEg2eyUpd2s0Qbsqhc4ZG33cI6VIFA5YgvS-wLrlb291JbtvJjMymwj5HcDj7YBnA/62fx62f",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          1,
		ItemName:        "Operation Broken Fang Case",
		ItemTransaction: 35,
		ItemPrice:       100000,
		ItemImage:       "https://community.akamai.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXU5A1PIYQNqhpOSV-fRPasw8rsQEl9Jg9SpIW1KgRr7OHFY28SvoyJl4iKm_vxPbzUhHgfuZEg2eyUpd2s0Qbsqhc4ZG33cI6VIFA5YgvS-wLrlb291JbtvJjMymwj5HcDj7YBnA/62fx62f",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          1,
		ItemName:        "Sticker | Natus Vincere | 2020 RMR",
		ItemTransaction: 16,
		ItemPrice:       600000,
		ItemImage:       "https://community.akamai.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXU5A1PIYQNqhpOSV-fRPasw8rsQEl9Jg9SpIW1KgRr7OHFY28SvoyJl4iKm_vxPbzUhHgfuZEg2eyUpd2s0Qbsqhc4ZG33cI6VIFA5YgvS-wLrlb291JbtvJjMymwj5HcDj7YBnA/62fx62f",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          2,
		ItemName:        "Glove Case",
		ItemTransaction: 456,
		ItemPrice:       500000,
		ItemImage:       "https://community.akamai.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXU5A1PIYQNqhpOSV-fRPasw8rsUFJ5KBFZv668FFY1naTMdzwTtNrukteIkqT2MO_Uwz5Q6cYhibyXo4rw2ALsrkRoYjuncNCLMlhpEV4XDTk/62fx62f",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          2,
		ItemName:        "Sticker | Natus Vincere (Holo) | 2020 RMR",
		ItemTransaction: 126,
		ItemPrice:       400000,
		ItemImage:       "https://community.akamai.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXU5A1PIYQNqhpOSV-fRPasw8rsUFJ5KBFZv668FFY1naTMdzwTtNrukteIkqT2MO_Uwz5Q6cYhibyXo4rw2ALsrkRoYjuncNCLMlhpEV4XDTk/62fx62f",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          2,
		ItemName:        "Fracture Case",
		ItemTransaction: 154,
		ItemPrice:       100000,
		ItemImage:       "https://community.akamai.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXU5A1PIYQNqhpOSV-fRPasw8rsUFJ5KBFZv668FFY1naTMdzwTtNrukteIkqT2MO_Uwz5Q6cYhibyXo4rw2ALsrkRoYjuncNCLMlhpEV4XDTk/62fx62f",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          2,
		ItemName:        "Sticker | Natus Vincere (Gold) | 2020 RMR",
		ItemTransaction: 112,
		ItemPrice:       200000,
		ItemImage:       "https://community.akamai.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXU5A1PIYQNqhpOSV-fRPasw8rsUFJ5KBFZv668FFY1naTMdzwTtNrukteIkqT2MO_Uwz5Q6cYhibyXo4rw2ALsrkRoYjuncNCLMlhpEV4XDTk/62fx62f",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          2,
		ItemName:        "Clutch Case",
		ItemTransaction: 54,
		ItemPrice:       700000,
		ItemImage:       "https://community.akamai.steamstatic.com/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXU5A1PIYQNqhpOSV-fRPasw8rsUFJ5KBFZv668FFY1naTMdzwTtNrukteIkqT2MO_Uwz5Q6cYhibyXo4rw2ALsrkRoYjuncNCLMlhpEV4XDTk/62fx62f",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	/* 10 item */

	db.Create(&GameItem{
		GameId:          3,
		ItemName:        "UMP-45 | Gold Bismuth",
		ItemTransaction: 125,
		ItemPrice:       400000,
		ItemImage:       "https://community.akamai.steamstatic.com/economy/image/IzMF03bk9WpSBq-S-ekoE33L-iLqGFHVaU25ZzQNQcXdA3g5gMEPvUZZEfSMJ6dESN8p_2SVTY7V2NwOx3QNmChCIzb02ClFcOt4PdmAyQj6pufQRHXxMGGRfyeBSl1rSOFZMjyMrDfz5unGQTjOSe94QA5WKaEN8zFAP9fTf0UrhthJr2CqqE1wHxEtL5QWIFjpnyVFM-8nnnYWdcpamiD3IpeLhQ1nYEVqUrrjA-mWZ9Wkwyg6HUQwCeYXZha0sNOx/62fx62f",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          3,
		ItemName:        "Sealed Graffiti | Rage Mode (Bazooka Pink)",
		ItemTransaction: 35,
		ItemPrice:       400000,
		ItemImage:       "https://community.akamai.steamstatic.com/economy/image/IzMF03bk9WpSBq-S-ekoE33L-iLqGFHVaU25ZzQNQcXdA3g5gMEPvUZZEfSMJ6dESN8p_2SVTY7V2NwOx3QNmChCIzb02ClFcOt4PdmAyQj6pufQRHXxMGGRfyeBSl1rSOFZMjyMrDfz5unGQTjOSe94QA5WKaEN8zFAP9fTf0UrhthJr2CqqE1wHxEtL5QWIFjpnyVFM-8nnnYWdcpamiD3IpeLhQ1nYEVqUrrjA-mWZ9Wkwyg6HUQwCeYXZha0sNOx/62fx62f",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          3,
		ItemName:        "FAMAS | Crypsis",
		ItemTransaction: 1336,
		ItemPrice:       500000,
		ItemImage:       "https://community.akamai.steamstatic.com/economy/image/IzMF03bk9WpSBq-S-ekoE33L-iLqGFHVaU25ZzQNQcXdA3g5gMEPvUZZEfSMJ6dESN8p_2SVTY7V2NwOx3QNmChCIzb02ClFcOt4PdmAyQj6pufQRHXxMGGRfyeBSl1rSOFZMjyMrDfz5unGQTjOSe94QA5WKaEN8zFAP9fTf0UrhthJr2CqqE1wHxEtL5QWIFjpnyVFM-8nnnYWdcpamiD3IpeLhQ1nYEVqUrrjA-mWZ9Wkwyg6HUQwCeYXZha0sNOx/62fx62f",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          3,
		ItemName:        "Bewitched (Foil)",
		ItemTransaction: 256,
		ItemPrice:       700000,
		ItemImage:       "https://community.akamai.steamstatic.com/economy/image/IzMF03bk9WpSBq-S-ekoE33L-iLqGFHVaU25ZzQNQcXdA3g5gMEPvUZZEfSMJ6dESN8p_2SVTY7V2NwOx3QNmChCIzb02ClFcOt4PdmAyQj6pufQRHXxMGGRfyeBSl1rSOFZMjyMrDfz5unGQTjOSe94QA5WKaEN8zFAP9fTf0UrhthJr2CqqE1wHxEtL5QWIFjpnyVFM-8nnnYWdcpamiD3IpeLhQ1nYEVqUrrjA-mWZ9Wkwyg6HUQwCeYXZha0sNOx/62fx62f",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          3,
		ItemName:        "RAVEN SHOTGUN | Cloaker Gold",
		ItemTransaction: 65,
		ItemPrice:       200000,
		ItemImage:       "https://community.akamai.steamstatic.com/economy/image/IzMF03bk9WpSBq-S-ekoE33L-iLqGFHVaU25ZzQNQcXdA3g5gMEPvUZZEfSMJ6dESN8p_2SVTY7V2NwOx3QNmChCIzb02ClFcOt4PdmAyQj6pufQRHXxMGGRfyeBSl1rSOFZMjyMrDfz5unGQTjOSe94QA5WKaEN8zFAP9fTf0UrhthJr2CqqE1wHxEtL5QWIFjpnyVFM-8nnnYWdcpamiD3IpeLhQ1nYEVqUrrjA-mWZ9Wkwyg6HUQwCeYXZha0sNOx/62fx62f",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          4,
		ItemName:        "The Steam Awards - 2020 Trading Card",
		ItemTransaction: 556,
		ItemPrice:       100000,
		ItemImage:       "https://community.akamai.steamstatic.com/economy/image/IzMF03bk9WpSBq-S-ekoE33L-iLqGFHVaU25ZzQNQcXdA3g5gMEPvUZZEfSMJ6dESN8p_2SVTY7V2NgPzXIJmD4QPivs0XEwSOp0PdiAyQj6pufQHievbWORdiWLGg9uRedXM2GMqjCj4rmWSz_IR-p-QA1RfqZX8mIdadfJYUUrhthJr2CqqE1wHxEtL5UUIFnjm3EVabt3mHEXIZ9bzST2dMfZ015nOhA4XL_vVu2XPtTzxCc6HUQwCeYXZviBbfYz/62fx62f",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          4,
		ItemName:        "MAG-7 | Rust Coat",
		ItemTransaction: 75,
		ItemPrice:       900000,
		ItemImage:       "https://community.akamai.steamstatic.com/economy/image/IzMF03bk9WpSBq-S-ekoE33L-iLqGFHVaU25ZzQNQcXdA3g5gMEPvUZZEfSMJ6dESN8p_2SVTY7V2NgPzXIJmD4QPivs0XEwSOp0PdiAyQj6pufQHievbWORdiWLGg9uRedXM2GMqjCj4rmWSz_IR-p-QA1RfqZX8mIdadfJYUUrhthJr2CqqE1wHxEtL5UUIFnjm3EVabt3mHEXIZ9bzST2dMfZ015nOhA4XL_vVu2XPtTzxCc6HUQwCeYXZviBbfYz/62fx62f",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          4,
		ItemName:        "Outstanding Visual Style",
		ItemTransaction: 77,
		ItemPrice:       100000,
		ItemImage:       "https://community.akamai.steamstatic.com/economy/image/IzMF03bk9WpSBq-S-ekoE33L-iLqGFHVaU25ZzQNQcXdA3g5gMEPvUZZEfSMJ6dESN8p_2SVTY7V2NgPzXIJmD4QPivs0XEwSOp0PdiAyQj6pufQHievbWORdiWLGg9uRedXM2GMqjCj4rmWSz_IR-p-QA1RfqZX8mIdadfJYUUrhthJr2CqqE1wHxEtL5UUIFnjm3EVabt3mHEXIZ9bzST2dMfZ015nOhA4XL_vVu2XPtTzxCc6HUQwCeYXZviBbfYz/62fx62f",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          4,
		ItemName:        "Sticker | Stone Scales",
		ItemTransaction: 34,
		ItemPrice:       150000,
		ItemImage:       "https://community.akamai.steamstatic.com/economy/image/IzMF03bk9WpSBq-S-ekoE33L-iLqGFHVaU25ZzQNQcXdA3g5gMEPvUZZEfSMJ6dESN8p_2SVTY7V2NgPzXIJmD4QPivs0XEwSOp0PdiAyQj6pufQHievbWORdiWLGg9uRedXM2GMqjCj4rmWSz_IR-p-QA1RfqZX8mIdadfJYUUrhthJr2CqqE1wHxEtL5UUIFnjm3EVabt3mHEXIZ9bzST2dMfZ015nOhA4XL_vVu2XPtTzxCc6HUQwCeYXZviBbfYz/62fx62f",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

	db.Create(&GameItem{
		GameId:          4,
		ItemName:        "Stone Gold Water",
		ItemTransaction: 21,
		ItemPrice:       200000,
		ItemImage:       "https://community.akamai.steamstatic.com/economy/image/IzMF03bk9WpSBq-S-ekoE33L-iLqGFHVaU25ZzQNQcXdA3g5gMEPvUZZEfSMJ6dESN8p_2SVTY7V2NgPzXIJmD4QPivs0XEwSOp0PdiAyQj6pufQHievbWORdiWLGg9uRedXM2GMqjCj4rmWSz_IR-p-QA1RfqZX8mIdadfJYUUrhthJr2CqqE1wHxEtL5UUIFnjm3EVabt3mHEXIZ9bzST2dMfZ015nOhA4XL_vVu2XPtTzxCc6HUQwCeYXZviBbfYz/62fx62f",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
		DeletedAt:       nil,
	})

}

func GetAllGameItem()([]GameItem, error){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var gameItem []GameItem
	db.Order("item_transaction desc").Find(&gameItem)

	return gameItem, nil
}

func GetItemGameById(id int) (GameItem, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var gameItem GameItem
	db.First(&gameItem, id)

	return gameItem, nil
}
