pokemon_details
-
"Pokemon ID" PK int 
"Name" string
"Description" string
"Japanese name" string
"Name etymology" string
"Image URL" string
"Species" string
"Height" int
"Weight" int
"Type 1" string
"Type 2" string
"Abilities" string
"EV Yield" int
"Catch rate" int
"Base Exp" int
"Growth rate" int
"Gender male %" float
"Gender female %" float
"Friendship" int
"Generation" int


pokemon_stats
-
"Pokemon ID" int FK >- pokemon_details."Pokemon ID" 
"Base HP" int
"Min. HP" int
"Max. HP" int
"Base Attack" int
"Min. Attack" int
"Max. Attack" int
"Base Defense" int
"Min. Defense" int
"Max. Defense" int
"Base Spd Attack" int
"Min. Spd Attack" int
"Max. Spd Attack" int
"Base Spd Defense" int
"Min. Spd Defense" int
"Max. Spd Defense" int
"Base Speed" int
"Min. Speed" int
"Max. Speed" int


pokemon_league
-
"League ID" int
"Pokemon ID" int FK >- pokemon_details."Pokemon ID" 
"Sprite URL" string
"League name" string
"League Region" string 

pokemon_region_details
-
"Region" int FK >- pokemon_league."League Region"
"Professor" string
"Region Villian" string 
"Starter Pokemon" string


pokemon_region_coordinates
-
"Region" string PK FK >- pokemon_region_details."Region"
"Coordinate X" int
"Coordinate Y" int