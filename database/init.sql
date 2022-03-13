use soda;

CREATE TABLE IF NOT EXISTS `client` (
  `id_client` int(11) NOT NULL AUTO_INCREMENT,
  `address` varchar(256) NOT NULL,
  `number` int(11) NOT NULL,
  `num_order` int(11) NOT NULL,
  `id_delivery` int(11) NOT NULL,
  `id_root` int(11) NOT NULL,
  `price_per_soda` decimal(10,0) NOT NULL,
  `price_per_box` decimal(10,0) NOT NULL,
  `debt` decimal(10,0) NOT NULL,
  PRIMARY KEY (`id_client`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `delivery`(
  `id_delivery` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(256) NOT NULL,
  PRIMARY KEY (`id_delivery`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `delivery_root` (
  `id_delivery` int(11) NOT NULL,
  `id_root` int(11) NOT NULL,
  `code` int(11) NOT NULL,
  PRIMARY KEY (`id_delivery`,`id_root`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO delivery(id_delivery, name)
VALUES  (1, 'Anselmo'),       
        (2, 'Fabian'),
        (3, 'Ariel'),
        (4, 'Juan'),
        (5, 'Mauricio'),
        (6, 'Jesus'),
        (7, 'Daniel');

INSERT INTO delivery_root (id_delivery, id_root, code)
VALUES  (1, 1, 1),
		(1, 2, 2),
		(1, 3, 3),
		(2, 4, 4),
		(2, 5, 5),
		(2, 6, 6),
		(3, 7, 7),
		(3, 8, 8),
		(3, 9, 9),
		(3, 12, 12),
		(4, 10, 10),
		(4, 11, 11),
		(7, 19, 19),
		(7, 20, 20),
		(7, 21, 21);
