CREATE TABLE users (
  user_id VARCHAR(20) NOT NULL,
  full_name VARCHAR(100) NOT NULL,
  password VARCHAR(50) NOT NULL,
  group_user VARCHAR(20) NOT NULL DEFAULT 'USER',
  balance INT DEFAULT 0,
  phone VARCHAR(20) NOT NULL,
  email VARCHAR(100) NOT NULL,
  isCustomer VARCHAR(1) DEFAULT 'N',
  isSeller VARCHAR(1) DEFAULT 'N',
  isShipper VARCHAR(1) DEFAULT 'N',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id)
);

CREATE TABLE user_addresses (
  address_id INT NOT NULL AUTO_INCREMENT,
  user_id VARCHAR(20) NOT NULL,
  address_line VARCHAR(500) NOT NULL,
  postal_code VARCHAR(10) NOT NULL,
  city VARCHAR(50) NOT NULL,
  phone VARCHAR(20) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (address_id)
);

CREATE TABLE hist_balance (
  user_id VARCHAR(20) NOT NULL,
  payment_id INT NOT NULL,
  amount INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE payments (
  payment_id INT NOT NULL AUTO_INCREMENT,
  payment_name VARCHAR(50) NOT NULL,
  isActive VARCHAR(1) DEFAULT 'N',
  created_by VARCHAR(20) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (payment_id)
);

-----
CREATE TABLE categories {
  category_id INT [pk, increment]
  category_name VARCHAR
  created_by VARCHAR
  created_at TIMESTAMP
  updated_at TIMESTAMP
}

CREATE TABLE products {
  product_id INT [pk, increment]
  product_name VARCHAR
  category_id VARCHAR
  price INT
  merchants_id VARCHAR
  created_at TIMESTAMP
  updated_at TIMESTAMP
}

CREATE TABLE merchants {
  merchants_id INT [pk, increment]
  product_name VARCHAR
  postal_code VARCHAR
  city VARCHAR
  phone VARCHAR
  user_id VARCHAR
  created_at TIMESTAMP
  updated_at TIMESTAMP
}

CREATE TABLE carts {
  cart_id INT [pk, increment]
  total_price INT
  user_id VARCHAR
  merchants_id VARCHAR
  created_at TIMESTAMP
}

CREATE TABLE cart_details {
  cart_id INT [pk]
  product_id VARCHAR [pk]
  product_name VARCHAR
  price INT
  quantity INT
  user_id VARCHAR
  merchants_id VARCHAR
  created_at TIMESTAMP
}

CREATE TABLE orders {
  order_id INT [pk, increment]
  total_price INT
  isPaid VARCHAR
  isShipped VARCHAR
  isReceived VARCHAR
  user_id VARCHAR
  cart_id VARCHAR
  expedition_id VARCHAR
  created_at TIMESTAMP
}

CREATE TABLE order_details {
  order_id INT [pk]
  product_id VARCHAR [pk]
  product_name VARCHAR
  price INT
  quantity INT
  user_id VARCHAR
  created_at TIMESTAMP
}

CREATE TABLE expeditions {
  expedition_id INT [pk, increment]
  user_id VARCHAR
  address_line VARCHAR
  postal_code VARCHAR
  city VARCHAR
  phone VARCHAR
  created_at TIMESTAMP
  updated_at TIMESTAMP
}