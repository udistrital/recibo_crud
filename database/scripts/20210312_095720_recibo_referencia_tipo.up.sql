ALTER TABLE recibo.recibo ADD COLUMN tercero_id INTEGER NOT NULL;

ALTER TABLE recibo.recibo ALTER COLUMN referencia SET DATA TYPE varchar(10);
