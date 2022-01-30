use 'soda';

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
