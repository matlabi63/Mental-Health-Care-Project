CREATE TABLE users (
    user_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    email_address VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role ENUM('End User', 'Doctor') NOT NULL,
    profile_picture VARCHAR(255),
    phone_number VARCHAR(20),
    address TEXT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE admin (
    admin_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    email_address VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20),
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE consultations (
    consultation_id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    doctor_id INT,
    consultation_date DATETIME NOT NULL,
    consultation_fee DECIMAL(10, 2) NOT NULL,
    status ENUM('Pending', 'Completed') DEFAULT 'Pending',
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (UserID) REFERENCES User(UserID),
    FOREIGN KEY (DoctorID) REFERENCES User(UserID)
);

CREATE TABLE payments (
    payment_id INT PRIMARY KEY AUTO_INCREMENT,
    consultation_id INT,
    user_id INT,
    amount DECIMAL(10, 2) NOT NULL,
    payment_date DATETIME NOT NULL,
    payment_status ENUM('Pending', 'Paid') DEFAULT 'Pending',
    FOREIGN KEY (ConsultationID) REFERENCES Consultation(ConsultationID),
    FOREIGN KEY (UserID) REFERENCES User(UserID)
);

CREATE TABLE complaints (
    complaint_id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    doctor_id INT,
    complaint_text TEXT NOT NULL,
    complaint_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    status ENUM('Open', 'Closed') DEFAULT 'Open',
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (UserID) REFERENCES User(UserID),
    FOREIGN KEY (DoctorID) REFERENCES User(UserID)
);

CREATE TABLE Advice (
    advice_id INT PRIMARY KEY AUTO_INCREMENT,
    complaint_id INT,
    doctor_id INT,
    advice_text TEXT NOT NULL,
    advice_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (ComplaintID) REFERENCES Complaint(ComplaintID),
    FOREIGN KEY (DoctorID) REFERENCES User(UserID)
);

CREATE TABLE chats (
    chat_id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    doctor_id INT,
    Message TEXT NOT NULL,
    Timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (UserID) REFERENCES User(UserID),
    FOREIGN KEY (DoctorID) REFERENCES User(UserID)
);

