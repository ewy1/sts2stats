package model

type PlayerStat struct {
	CurrentGold int `json:"current_gold"`
	CurrentHp   int `json:"current_hp"`
	DamageTaken int `json:"damage_taken"`
	GoldGained  int `json:"gold_gained"`
	GoldLost    int `json:"gold_lost"`
	GoldSpent   int `json:"gold_spent"`
	GoldStolen  int `json:"gold_stolen"`
	HpHealed    int `json:"hp_healed"`
	MaxHp       int `json:"max_hp"`
	MaxHpGained int `json:"max_hp_gained"`
	MaxHpLost   int `json:"max_hp_lost"`
	PlayerID    int `json:"player_id"`
}

type RunSave struct {
	Acts              []string `json:"acts"`
	Ascension         int      `json:"ascension"`
	BuildID           string   `json:"build_id"`
	GameMode          string   `json:"game_mode"`
	KilledByEncounter string   `json:"killed_by_encounter"`
	KilledByEvent     string   `json:"killed_by_event"`
	MapPointHistory   [][]struct {
		MapPointType string `json:"map_point_type"`
		PlayerStats  []struct {
			PlayerStat
			CardChoices []struct {
				Card struct {
					Id                  string `json:"id"`
					CurrentUpgradeLevel int    `json:"current_upgrade_level"`
				}
				WasPicked bool `json:"was_picked"`
			} `json:"card_choices"`
			AncientChoice []struct {
				TextKey string `json:"TextKey"`
				Title   struct {
					Key   string `json:"key"`
					Table string `json:"table"`
				} `json:"title"`
				WasChosen bool `json:"was_chosen"`
			} `json:"ancient_choice"`
			CardsTransformed []struct {
				FinalCard struct {
					FloorAddedToDeck int    `json:"floor_added_to_deck"`
					ID               string `json:"id"`
				} `json:"final_card"`
				OriginalCard struct {
					FloorAddedToDeck int    `json:"floor_added_to_deck"`
					ID               string `json:"id"`
				} `json:"original_card"`
			} `json:"cards_transformed"`
			EventChoices []struct {
				Title struct {
					Key   string `json:"key"`
					Table string `json:"table"`
				} `json:"title"`
			} `json:"event_choices"`
			GoldGained   int `json:"gold_gained"`
			GoldLost     int `json:"gold_lost"`
			GoldSpent    int `json:"gold_spent"`
			GoldStolen   int `json:"gold_stolen"`
			HpHealed     int `json:"hp_healed"`
			MaxHp        int `json:"max_hp"`
			MaxHpGained  int `json:"max_hp_gained"`
			MaxHpLost    int `json:"max_hp_lost"`
			PlayerID     int `json:"player_id"`
			RelicChoices []struct {
				Choice    string `json:"choice"`
				WasPicked bool   `json:"was_picked"`
			} `json:"relic_choices"`
		} `json:"player_stats"`
		Rooms []struct {
			ModelID    string `json:"model_id"`
			RoomType   string `json:"room_type"`
			TurnsTaken int    `json:"turns_taken"`
		} `json:"rooms"`
	} `json:"map_point_history"`
	Modifiers    []interface{} `json:"modifiers"`
	PlatformType string        `json:"platform_type"`
	Players      []struct {
		Character string `json:"character"`
		Deck      []struct {
			FloorAddedToDeck int    `json:"floor_added_to_deck"`
			ID               string `json:"id"`
			Enchantment      struct {
				Amount int    `json:"amount"`
				ID     string `json:"id"`
			} `json:"enchantment,omitempty"`
		} `json:"deck"`
		ID                 int           `json:"id"`
		MaxPotionSlotCount int           `json:"max_potion_slot_count"`
		Potions            []interface{} `json:"potions"`
		Relics             []struct {
			FloorAddedToDeck int    `json:"floor_added_to_deck"`
			ID               string `json:"id"`
		} `json:"relics"`
	} `json:"players"`
	RunTime       int    `json:"run_time"`
	SchemaVersion int    `json:"schema_version"`
	Seed          string `json:"seed"`
	StartTime     int    `json:"start_time"`
	WasAbandoned  bool   `json:"was_abandoned"`
	Win           bool   `json:"win"`
}
