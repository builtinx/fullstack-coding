CREATE DATABASE fullstack;
USE fullstack;

CREATE TABLE companies (
	id INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
	name VARCHAR(256) NOT NULL DEFAULT '',
	mission VARCHAR(256) NOT NULL DEFAULT '',
	PRIMARY KEY (id),
	UNIQUE KEY name (name)
) ENGINE = InnoDB CHARSET=utf8mb4;

INSERT INTO companies (name, mission) VALUES ('Secondary', 'We program the open-source so you can bypass your brooklyn');
INSERT INTO companies (name, mission) VALUES ('Cohesive', 'We program the viral so you can hack your gluten-free');
INSERT INTO companies (name, mission) VALUES ('Approach', 'We calculate the channels so you can program your franzen');
INSERT INTO companies (name, mission) VALUES ('Mobile', 'We generate the synthesize so you can hack your banh mi');
INSERT INTO companies (name, mission) VALUES ('Pricing Structure', 'We override the exploit so you can program your paleo');
INSERT INTO companies (name, mission) VALUES ('Adapter', 'We input the wireless so you can reboot your food truck');
INSERT INTO companies (name, mission) VALUES ('Internet Solution', 'We synthesize the reinvent so you can generate your selfies');
INSERT INTO companies (name, mission) VALUES ('Encryption', 'We index the niches so you can program your whatever');
INSERT INTO companies (name, mission) VALUES ('Focused', 'We navigate the bandwidth so you can quantify your tacos');
INSERT INTO companies (name, mission) VALUES ('De-engineered', 'We back up the holistic so you can generate your umami');
