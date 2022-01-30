use 'soda';

CREATE TABLE `client` (
  `id_client` int(11) NOT NULL AUTO_INCREMENT,
  `address` varchar(256) NOT NULL,
  `number` int(11) NOT NULL,
  `order` int(11) NOT NULL,
  `id_delivery` int(11) NOT NULL,
  `id_root` int(11) NOT NULL,
  `price_per_soda` decimal(10,0) NOT NULL,
  `price_per_box` decimal(10,0) NOT NULL,
  PRIMARY KEY (`id_client`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

CREATE TABLE `delivery` (
  `id_delivery` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(256) NOT NULL,
  PRIMARY KEY (`id_delivery`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;

CREATE TABLE `delivery_root` (
  `id_delivery` int(11) NOT NULL,
  `id_root` int(11) NOT NULL,
  `code` int(11) NOT NULL,
  PRIMARY KEY (`id_delivery`,`id_root`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

