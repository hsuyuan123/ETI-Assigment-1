CREATE database my_db;
CREATE USER 'user1'@'localhost' IDENTIFIED BY 'password';
GRANT ALL ON *.* TO 'user1'@'localhost';

USE my_db;

CREATE TABLE User (UserId varchar (5) NOT NULL PRIMARY KEY, FirstName VARCHAR(30), LastName VARCHAR(30), MobileNo INT, Email VARCHAR(30), accstatus varchar(10)); 

CREATE TABLE Driver (DriverId varchar (5) NOT NULL PRIMARY KEY, FirstName VARCHAR(30), LastName VARCHAR(30), MobileNo INT, Email VARCHAR(30), VehicleNo VARCHAR(30), LicenseNo VARCHAR(30), accstatus VARCHAR(10)); 

CREATE TABLE Trip (TripId varchar (5) NOT NULL PRIMARY KEY, DriverID varchar (5), PickupLoc VARCHAR(50), AltPickupLoc VARCHAR(50), StartTravelTime DATETIME, DestAddress VARCHAR(50), NoOfPsgr INT, UserId varchar (5)); 

