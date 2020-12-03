-- Create schema for project --
CREATE SCHEMA SeatSafe;
USE SeatSafe;

-- Create tables -- 
CREATE TABLE Event (
    EventId bigint NOT NULL AUTO_INCREMENT PRIMARY KEY,
    PublicId varchar(36) NOT NULL,
    PrivateId varchar(36) NOT NULL,
    EventName varchar(100) NOT NULL,
    ContactEmail varchar(100) NOT NULL,
    PublicallyListed boolean NOT NULL,
    ImageUrl varchar(150) DEFAULT NULL,
    INDEX (PublicId),
    INDEX (PrivateId)
);
ALTER TABLE Event AUTO_INCREMENT = 1000000000;

CREATE TABLE SpotGroup (
    SpotGroupId bigint NOT NULL AUTO_INCREMENT PRIMARY KEY,
    EventId bigint NOT NULL,
    Name varchar(100) NOT NULL,
    FOREIGN KEY (EventId)
        REFERENCES Event(EventId)
        ON DELETE CASCADE,
    INDEX (EventId) 
);
ALTER TABLE SpotGroup AUTO_INCREMENT = 2000000000;

CREATE TABLE Reservation (
    ReservationId bigint NOT NULL AUTO_INCREMENT PRIMARY KEY,
    PrivateId varchar(36) NOT NULL,
    Email varchar(100) NOT NULL,
    Name varchar(100) NOT NULL,
    EventId bigint NOT NULL,
    FOREIGN KEY (EventId)
        REFERENCES Event(EventId)
        ON DELETE CASCADE
);
ALTER TABLE Reservation AUTO_INCREMENT = 3000000000;

CREATE TABLE Spot (
    SpotId bigint NOT NULL AUTO_INCREMENT PRIMARY KEY,
    SpotGroupId bigint NOT NULL,
    ReservationId bigint DEFAULT NULL,
    FOREIGN KEY (ReservationId)
        REFERENCES Reservation(ReservationId)
        ON DELETE SET NULL,
    FOREIGN KEY (SpotGroupId)
        REFERENCES SpotGroup(SpotGroupId)
        ON DELETE CASCADE
);
ALTER TABLE Spot AUTO_INCREMENT = 4000000000;