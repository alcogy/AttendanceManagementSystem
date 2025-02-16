CREATE TABLE `User` (
  `id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `auth` tinyint NOT NULL,
  `active` tinyint NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE `Attendance` (
  `id` varchar(255) NOT NULL,
  `date` date NOT NULL,
  `inTime` time,
  `outTime` time,
  `breakHour` time,
  `userId` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`userId`) REFERENCES `User`(`id`),
  UNIQUE uniqueDateUser(`date`, `userId`),
  INDEX indexUser(`userId`)
);

CREATE TABLE `BreakTime` (
  `attendanceId` varchar(255),
  `inTime` time,
  `outTime` time,
  FOREIGN KEY (`attendanceId`) REFERENCES `User`(`id`),
  INDEX indexAttendance(`attendanceId`)
);