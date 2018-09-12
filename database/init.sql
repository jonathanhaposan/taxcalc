CREATE TABLE tax_code
(
  code_id integer NOT NULL,
  description character varying(64),
  CONSTRAINT primary_key PRIMARY KEY (code_id)
);

CREATE TABLE bill
(
  bill_id SERIAL,
  product_name character varying(64),
  original_price integer,
  tax_amount numeric,
  total_amount numeric,
  tax_code_id integer,
  CONSTRAINT pk PRIMARY KEY (bill_id),
  CONSTRAINT fk1 FOREIGN KEY (tax_code_id)
      REFERENCES tax_code (code_id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
);

INSERT INTO tax_code(code_id, description) VALUES (1, 'Food');
INSERT INTO tax_code(code_id, description) VALUES (2, 'Tobacco');
INSERT INTO tax_code(code_id, description) VALUES (3, 'Entertainment');