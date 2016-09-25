# Document, Don't Create

go get github.com/gorilla/mux
go get github.com/dgrijalva/jwt-go
go get github.com/go-sql-driver/mysql

```
CREATE TABLE `pokemon` (
  `pokedexID` int(11) DEFAULT NULL,
  `name` varchar(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8
```

```
CREATE TABLE `pokemon_type` (
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `pokemonID` int(11) NOT NULL,
  `type_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8
```

```
CREATE TABLE `types` (
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8
```

