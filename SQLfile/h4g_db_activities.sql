-- MySQL dump 10.13  Distrib 8.0.32, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: h4g_db
-- ------------------------------------------------------
-- Server version	8.0.32

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `activities`
--

DROP TABLE IF EXISTS `activities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `activities` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  `desc` longtext,
  `time` longtext,
  `location` longtext,
  `creator_type` bigint unsigned DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_activities_user` (`user_id`),
  CONSTRAINT `fk_activities_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `activities`
--

LOCK TABLES `activities` WRITE;
/*!40000 ALTER TABLE `activities` DISABLE KEYS */;
INSERT INTO `activities` VALUES (1,'2023-02-11 09:34:14.993','Teaching table tennis!','I am a professional table tennis player and I have some free time! I would like to teach my skills to people for free. Everyone is welcome, I am extremely patient! We can all be friends!','3 March 2023 at 4.00P.M.','NUS UTSH Table Tennis room',1,1),(8,'2023-02-11 10:12:43.765','Please teach me paralympic basketball','Hi, my name is Sam! I always wanted to play basketball but I am wheelchair bound. If someone knows how to play paralympic basketball, please teach me!','28 Feb 2023, 10.00 A.M.','Simei MRT basketball court',0,2),(9,'2023-02-11 10:19:47.494','Walk with me around the CBD area!','Hi, I am looking for people to walk with me around the CBD area for exercise. If you have difficulties getting around, I would love to help you. Hope to see you soon!','13 Feb 2023 at 6.00P.M.','CBD area (Marina Bay Sands)',1,3),(10,'2023-02-11 10:22:41.311','Teaching Braille','Join me on a journey of empowerment and learning! As a passionate advocate for the visually impaired, I am offering free Braille lessons to anyone who wants to improve their literacy skills and gain independence. Whether you are a beginner or have some experience, I will guide you through the basics of Braille, including letter recognition, number systems, and punctuation. With hands-on exercises, interactive activities, and individualized attention, I guarantee that you will leave with a strong foundation in this essential system of communication. Don\'t miss this opportunity to enhance your life and make a difference for others.','Every Wednesday from 15 Feb 2023 at 10.00 A.M.','Canberra Block 31, #10-341',0,2),(11,'2023-02-11 10:27:28.512','Lorem Ipsum','What is Lorem Ipsum?\nLorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.','11 February 2023, any time','NUS Tembusu College',1,3);
/*!40000 ALTER TABLE `activities` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-02-11 11:05:00
