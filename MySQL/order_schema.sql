CREATE DATABASE ORDERS;

Use ORDERS;

-- Create order table
CREATE TABLE `orders` (
    id INT AUTO_INCREMENT,
	start_latitude VARCHAR(255) NOT NULL,
	start_longtitude VARCHAR(255) NOT NULL,
	end_latitude VARCHAR(255) NOT NULL,
	end_longtitude VARCHAR(255) NOT NULL,
	status VARCHAR(10) NULL,
	distance INT NULL,
	pub_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (id)
);

INSERT INTO orders (
	start_latitude, start_longtitude, end_latitude, end_longtitude, distance, status
	) VALUES ('33333.5', '444444.3', '666666.4', '77777.8', 5000, "UNASIGNED");