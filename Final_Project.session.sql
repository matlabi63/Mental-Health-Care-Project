CREATE TABLE users (
    user_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    email_address VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role ENUM('End User', 'Doctor') NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


CREATE TABLE recommendations (
    id INT PRIMARY KEY AUTO_INCREMENT,
    advice TEXT NOT NULL,
    doctor_id INT,
    patient_id INT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (doctor_id) REFERENCES Users(UserID),
    FOREIGN KEY (patient_id) REFERENCES Users(UserID)
);


CREATE TABLE complaints (
    id INT PRIMARY KEY AUTO_INCREMENT,
    doctor_id INT,
    description TEXT NOT NULL,
    submission DATETIME DEFAULT CURRENT_TIMESTAMP,
    status ENUM('Open', 'Closed') DEFAULT 'Open',
    patient_id INT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (UserID) REFERENCES Users(UserID),
    FOREIGN KEY (DoctorID) REFERENCES Users(UserID)
);

CREATE TABLE informations (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255),
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE admins (
    user_id INT PRIMARY KEY,
    manage_users BOOLEAN,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
)

CREATE TABLE doctors (
    user_id INT PRIMARY KEY,
    special_users VARCHAR(255),
    available_users VARCHAR(255)
    FOREIGN KEY (user_id) REFERENCES users(user_id)
)

CREATE TABLE patients (
    user_id INT PRIMARY KEY,
    mental_status VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
)